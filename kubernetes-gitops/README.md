Ссылка на репозиторий GitLab: https://gitlab.com/wbe7/microservices-demo

ts=2020-09-14T12:11:33.627298982Z caller=loop.go:134 component=sync-loop event=refreshed url=ssh://git@gitlab.com/wbe7/microservices-demo.git branch=master HEAD=8c385c012e3e0dd012348568ff5bbbe1c9d218f0
ts=2020-09-14T12:11:40.059066826Z caller=sync.go:61 component=daemon info="trying to sync git changes to the cluster" old=8c385c012e3e0dd012348568ff5bbbe1c9d218f0 new=8c385c012e3e0dd012348568ff5bbbe1c9d218f0
ts=2020-09-14T12:11:41.275238831Z caller=sync.go:540 method=Sync cmd=apply args= count=2
ts=2020-09-14T12:11:41.529637222Z caller=sync.go:606 method=Sync cmd="kubectl apply -f -" took=254.321455ms err=null output="namespace/microservices-demo unchanged\nhelmrelease.helm.fluxcd.io/frontend unchanged"
ts=2020-09-14T12:12:33.625368818Z caller=images.go:17 component=sync-loop msg="polling for new images for automated workloads"
ts=2020-09-14T12:12:36.06362204Z caller=loop.go:134 component=sync-loop event=refreshed url=ssh://git@gitlab.com/wbe7/microservices-demo.git branch=master HEAD=8b5ec506b7664c2f70882ff3ca00e8a3023c19d9
ts=2020-09-14T12:12:36.067764531Z caller=sync.go:61 component=daemon info="trying to sync git changes to the cluster" old=8c385c012e3e0dd012348568ff5bbbe1c9d218f0 new=8b5ec506b7664c2f70882ff3ca00e8a3023c19d9
ts=2020-09-14T12:12:37.237494604Z caller=sync.go:540 method=Sync cmd=apply args= count=2
ts=2020-09-14T12:12:37.55712082Z caller=sync.go:606 method=Sync cmd="kubectl apply -f -" took=319.551258ms err=null output="namespace/microservices-demo unchanged\nhelmrelease.helm.fluxcd.io/frontend unchanged"
ts=2020-09-14T12:12:37.559513418Z caller=daemon.go:701 component=daemon event="Sync: 8b5ec50, no workloads changed" logupstream=false
ts=2020-09-14T12:12:40.228279383Z caller=loop.go:236 component=sync-loop state="tag flux-sync" old=8c385c012e3e0dd012348568ff5bbbe1c9d218f0 new=8b5ec506b7664c2f70882ff3ca00e8a3023c19d9
ts=2020-09-14T12:12:42.857248552Z caller=loop.go:134 component=sync-loop event=refreshed url=ssh://git@gitlab.com/wbe7/microservices-demo.git branch=master HEAD=8b5ec506b7664c2f70882ff3ca00e8a3023c19d9

NAME       STATUS      WEIGHT   LASTTRANSITIONTIME
frontend   Succeeded   0        2020-09-15T18:35:32Z


Events:
  Type     Reason  Age                  From     Message
  ----     ------  ----                 ----     -------
  Warning  Synced  23m                  flagger  frontend-primary.microservices-demo not ready: waiting for rollout to finish: observed deployment generation less then desired generation
  Normal   Synced  22m (x2 over 23m)    flagger  all the metrics providers are available!
  Normal   Synced  22m                  flagger  Initialization done! frontend.microservices-demo
  Normal   Synced  13m                  flagger  (combined from similar events): Promotion completed! Scaling down frontend.microservices-demo
  Normal   Synced  7m31s (x2 over 19m)  flagger  New revision detected! Scaling up frontend.microservices-demo
  Normal   Synced  6m31s (x2 over 18m)  flagger  Advance frontend.microservices-demo canary weight 10
  Normal   Synced  6m31s (x2 over 18m)  flagger  Starting canary analysis for frontend.microservices-demo
  Normal   Synced  5m31s (x2 over 17m)  flagger  Advance frontend.microservices-demo canary weight 20
  Normal   Synced  4m30s (x2 over 16m)  flagger  Advance frontend.microservices-demo canary weight 30
  Normal   Synced  3m31s (x2 over 15m)  flagger  Copying frontend.microservices-demo template spec to frontend-primary.microservices-demo
  Normal   Synced  2m30s (x2 over 14m)  flagger  Routing all traffic to primary
  Normal   Synced  89s                  flagger  Promotion completed! Scaling down frontend.microservices-demo
