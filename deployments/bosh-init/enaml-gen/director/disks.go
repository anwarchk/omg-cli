package director 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Disks struct {

	/*CleanupSchedule - Descr: RufusScheduler cron formatted schedule for cleanup of orphaned disks and orphaned snapshots Default: 0 0,30 * * * * UTC
*/
	CleanupSchedule interface{} `yaml:"cleanup_schedule,omitempty"`

	/*MaxOrphanedAgeInDays - Descr: Days to keep orphaned disks and orhaned snapshots before cleanup Default: 5
*/
	MaxOrphanedAgeInDays interface{} `yaml:"max_orphaned_age_in_days,omitempty"`

}