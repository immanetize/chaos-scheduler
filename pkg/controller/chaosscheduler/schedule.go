package chaosscheduler

import (
	"errors"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ref "k8s.io/client-go/tools/reference"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

  ctrl "sigs.k8s.io/controller-runtime"
  operatorv1 "github.com/litmuschaos/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
  schedulerv1alpha1 "github.com/litmuschaos/chaos-scheduler/api/v1alpha1"
  chaosTypes "github.com/litmuschaos/chaos-scheduler/controllers/types"
  
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func schedule(schedulerReconcile *reconcileScheduler, scheduler *chaosTypes.SchedulerInfo) (ctrl.Result, error) {

	if scheduler.Instance.Spec.Schedule.Now == true {
		schedulerReconcile.reqLogger.Info("Current scheduler type derived is ", "schedulerType", "now")
		return schedulerReconcile.createForNowAndOnce(scheduler)

	} else if scheduler.Instance.Spec.Schedule.Once != nil {
		schedulerReconcile.reqLogger.Info("Current scheduler type derived is ", "schedulerType", "once")
		scheduleTime := time.Now()
		startDuration := scheduler.Instance.Spec.Schedule.Once.ExecutionTime.Local().Sub(scheduleTime)

		if startDuration.Seconds() < 0 {
			if scheduler.Instance.Spec.Schedule.Once.ExecutionTime.Time.Before(scheduleTime) {
				return schedulerReconcile.createForNowAndOnce(scheduler)
			}
			schedulerReconcile.reqLogger.Info("ExecutionTime elapsed before the schedule creation")
		}
		schedulerReconcile.reqLogger.Info("Time left to schedule the engine", "Duration", startDuration)
		return ctrl.Result{RequeueAfter: startDuration}, nil

	} else if scheduler.Instance.Spec.Schedule.Repeat != nil {
		schedulerReconcile.reqLogger.Info("Current scheduler type derived is ", "schedulerType", "repeat")
		/* StartDuration is the duration between current time
		 * and the scheduled time to start the chaos which is
		 * being used by reconciler to reque this resource after
		 * that much duration
		 * Chaos is being started 1 min before the scheduled time
		 */

		var startTime *metav1.Time
		if scheduler.Instance.Spec.Schedule.Repeat.TimeRange != nil {
			startTime = scheduler.Instance.Spec.Schedule.Repeat.TimeRange.StartTime
		}

		if startTime == nil {
			startTime = &scheduler.Instance.CreationTimestamp
		}

		scheduleTime := time.Now()
		startDuration := startTime.Local().Sub(scheduleTime)

		if startDuration.Seconds() < 0 {
			return schedulerReconcile.createEngineRepeat(scheduler)
		}
		schedulerReconcile.reqLogger.Info("Time left to schedule the engine", "Duration", startDuration)
		return ctrl.Result{RequeueAfter: startDuration}, nil

	}

	return ctrl.Result{}, errors.New("ScheduleType should be one of ('now', 'once', 'repeat')")
}

func (r *ChaosscheduleReconciler) getRef(object runtime.Object) (*corev1.ObjectReference, error) {
	return ref.GetReference(r.scheme, object)
}

// getEngineFromTemplate makes an Engine from a Schedule
func getEngineFromTemplate(cs *chaosTypes.SchedulerInfo) *operatorV1.ChaosEngine {

	labels := map[string]string{
		"app":      "chaos-engine",
		"chaosUID": string(cs.Instance.UID),
	}

	engine := &operatorV1.ChaosEngine{}

	ownerReferences := make([]metav1.OwnerReference, 0)
	ownerReferences = append(ownerReferences, *metav1.NewControllerRef(cs.Instance, schedulerV1.SchemeGroupVersion.WithKind("ChaosSchedule")))
	engine.SetOwnerReferences(ownerReferences)
	engine.SetLabels(labels)

	engine.Spec = cs.Instance.Spec.EngineTemplateSpec
	engine.Spec.EngineState = operatorV1.EngineStateActive

	return engine
}
