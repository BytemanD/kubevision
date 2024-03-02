package apiv1

import "github.com/gogf/gf/v2/frame/g"

type CronJobsListReq struct {
	g.Meta `path:"/cronjobs" tags:"CronJobs" method:"get"`
}
type CronJobsListRes struct {
	g.Meta `mime:"application/json" example:"{\"cronjobs\":[]}"`
}
type CronJobsPostReq struct {
	g.Meta `path:"/cronjobs" tags:"CronJobs" method:"post"`
}
type CronJobsPostRes struct {
	g.Meta `mime:"application/json"`
}
type CronJobsDeleteReq struct {
	g.Meta `path:"/cronjobs" tags:"CronJobs" method:"delete"`
}
type CronJobsDeleteRes struct {
	g.Meta `mime:"application/json"`
}
