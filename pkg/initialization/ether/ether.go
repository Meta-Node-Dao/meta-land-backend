/**
 * @Author: Sun
 * @Description:
 * @File:  bounty
 * @Version: 1.0.0
 * @Date: 2022/7/1 20:38
 */

package ether

import (
	"ceres/pkg/service/ether"
	"github.com/gotomicro/ego/task/ecron"
)

func Init() ecron.Ecron {
	cron := ether.GetAllContractAddresses()
	return cron
}
