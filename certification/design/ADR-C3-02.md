# ADR-C3-02: CCE with horizontal pod autoscaling on CPU

## Status

Proposed

## Context

Section 3 of the design note needs Appointments and Clinical to grow from 2 to 10 instances when average CPU passes 70 %, and NFR-003 expects new capacity within 2 minutes of a spike. Section 2 (Performance efficiency and Operations rows) shows that no service exports metrics today, so nothing can trigger scaling or show whether the P95 target of 300 ms holds. The services are already stateless and packaged as Docker images (NFR-007), but they only run through Docker Compose on one machine, which cannot add instances on its own.

## Decision

Run the gateway and the five services as Deployments in one CCE cluster, with images pulled from SWR. Appointments and Clinical get a Horizontal Pod Autoscaler with a target of 70 % average CPU, a minimum of 2 pods and a maximum of 10, and the other services start with a fixed 2 pods. The CCE node pool keeps one spare node per zone so new pods start without waiting for a new ECS. AOM collects the CPU and RED metrics that drive the autoscaler and the alarms, and readiness probes on `/health/ready` keep traffic away from pods that are not ready.

## Alternatives considered

- Scale ECS instances with Auto Scaling groups and Docker Compose on each VM: rejected because scaling whole VMs is slower than scheduling pods and the team would need its own image rollout on every VM.
- Keep a fixed number of instances sized for the peak: rejected because it pays for 10 instances per service all day while the clinic has clear quiet hours.

## Consequences

- Positive: capacity follows the real load within the 2-minute target, and the same image is promoted between namespaces, as NFR-007 requires.
- Negative: the team must learn and operate Kubernetes manifests, and CPU is only an indirect signal; a slow MongoDB query can raise latency without raising CPU.
