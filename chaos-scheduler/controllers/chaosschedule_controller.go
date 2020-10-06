/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cachev1alpha1 "github.com/litmuschaos/chaos-scheduler/api/v1alpha1"
)

// ChaosscheduleReconciler reconciles a Chaosschedule object
type ChaosscheduleReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// reconcileScheduler contains details of reconcileScheduler
type reconcileScheduler struct {
	r         *ChaosscheduleReconciler
	reqLogger logr.Logger
}

// these verbs and resource types are very generous - this is an admittedly ignorant configuration during migration to new operator-sdk.
// code or runtime review should be performed to limit to least required privileges
// +kubebuilder:rbac:groups=cache.litmuschaos.io,resources=chaosschedules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.litmuschaos.io,resources=chaosschedules/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=batch/v1,resources=cronjobs;jobs,verbs=create;delete;list;patch;update;watch
// +kubebuilder:festsrbac:groups=apps/v1,resources=daemonsets;deployments;replicasets;statefulsets,verbs=create;delete;list;patch;update;watch

func (r *ChaosscheduleReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("chaosschedule", req.NamespacedName)

	// your logic here
        scheduler, err := r.getChaosSchedulerInstance(ctrl)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	schedulerReconcile := &reconcileScheduler{
		r:         r,
		reqLogger: reqLogger,
	}

	switch scheduler.Instance.Spec.ScheduleState {
	case "", schedulerV1.StateActive:
		{
			return schedulerReconcile.reconcileForCreationAndRunning(scheduler)
		}
	case schedulerV1.StateCompleted:
		{
			if !checkScheduleStatus(scheduler, schedulerV1.StatusCompleted) {
				return schedulerReconcile.reconcileForComplete(scheduler, request)
			}
		}
	case schedulerV1.StateHalted:
		{
			if !checkScheduleStatus(scheduler, schedulerV1.StatusHalted) {
				return schedulerReconcile.reconcileForHalt(scheduler, request)
			}
		}
	}
	return reconcile.Result{}, nil
}

func (r *ChaosscheduleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.Chaosschedule{}).
                Owns(&operatorV1.ChaosEngine{}).
		Complete(r)
}
func (schedulerReconcile *reconcileScheduler) reconcileForComplete(cs *chaosTypes.SchedulerInfo, request reconcile.Request) (reconcile.Result, error) {

	if len(cs.Instance.Status.Active) != 0 {
		errUpdate := schedulerReconcile.r.updateActiveStatus(cs)
		if errUpdate != nil {
			return reconcile.Result{}, errUpdate
		}
		return reconcile.Result{}, nil
	}

	opts := client.UpdateOptions{}
	cs.Instance.Status.Schedule.Status = schedulerV1.StatusCompleted
	cs.Instance.Status.Schedule.EndTime = &metav1.Time{Time: time.Now()}
	if err := schedulerReconcile.r.client.Update(context.TODO(), cs.Instance, &opts); err != nil {
		schedulerReconcile.r.recorder.Eventf(cs.Instance, corev1.EventTypeWarning, "ScheduleCompleted", "Cannot update status as completed")
		return reconcile.Result{}, fmt.Errorf("Unable to update chaosSchedule for status completed, due to error: %v", err)
	}
	schedulerReconcile.r.recorder.Eventf(cs.Instance, corev1.EventTypeNormal, "ScheduleCompleted", "Schedule completed successfully")
	return reconcile.Result{}, nil
}

func (schedulerReconcile *reconcileScheduler) reconcileForCreationAndRunning(cs *chaosTypes.SchedulerInfo) (reconcile.Result, error) {

	reconcileRes, err := schedule(schedulerReconcile, cs)
	if err != nil {
		return reconcile.Result{}, err
	}

	return reconcileRes, nil
}

func checkScheduleStatus(cs *chaosTypes.SchedulerInfo, status schedulerV1.ChaosStatus) bool {
	return cs.Instance.Status.Schedule.Status == status
}

// Fetch the ChaosScheduler instance
func (r *ReconcileChaosScheduler) getChaosSchedulerInstance(request ctrl.Request) (*chaosTypes.SchedulerInfo, error) {
	instance := &schedulerV1.ChaosSchedule{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		// Error reading the object - requeue the request.
		return nil, err
	}
	scheduler := &chaosTypes.SchedulerInfo{
		Instance: instance,
	}
	return scheduler, nil
}
