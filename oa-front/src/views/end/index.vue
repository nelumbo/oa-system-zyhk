<template>
    <div v-if="(base.model.ice == 0 || base.model.ice == 1) && user().my.pids.includes('75')">
        <el-divider content-position="left">
            <h2>结算</h2>
        </el-divider>
        <settlement :ice="base.model.ice" />
    </div>

    <div style="margin-top: 30px;"
        v-if="base.model.ice == 2 && (base.model.employee.office.isSetSubmit == -1 || base.model.employee.office.isSetSubmit == 0 || base.model.employee.office.isSetSubmit == 1) && user().my.pids.includes('89')">
        <el-divider content-position="left">
            <h2>提交提成分配方案</h2>
        </el-divider>
        <h2 v-if="base.model.employee.office.isSetSubmit == -1">分配方案被否决分发！！请重新提交分配方案</h2>
        <dispense :employee="base.model.employee" />
    </div>

    <div style="margin-top: 30px;" v-if="base.model.ice == 2 && user().my.pids.includes('75')">
        <el-divider content-position="left">
            <h2>办事处提成分配审批</h2>
        </el-divider>
        <approve />
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { reactive, onBeforeMount } from 'vue'
import { queryMe } from "@/api/my"
import { iceGet } from "@/api/settlement"

import settlement from './Settlement.vue'
import dispense from './Dispense.vue'
import approve from './Approve.vue'

const base = reactive({
    model: {
        ice: 2,
        employee: {
            office: {
                isSetSubmit: 2
            }
        },
    }
})

onBeforeMount(() => {
    iceGet().then((res) => {
        if (res.status == 1) {
            base.model.ice = res.data
        }
    })
    queryMe().then((res) => {
        if (res.status == 1) {
            base.model.employee = res.data
        }
    })
})
</script>