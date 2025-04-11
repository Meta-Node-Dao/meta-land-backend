/**
 * @Author: Sun
 * @Description:
 * @File:  time
 * @Version: 1.0.0
 * @Date: 2022/6/30 12:08
 */

package tool

import "time"

func ParseTimeString2Time(expireIn string) time.Time {
	timeTemplate := "2006-01-02 15:04:05"
	tamp, _ := time.ParseInLocation(timeTemplate, expireIn, time.Local)
	return tamp
}
