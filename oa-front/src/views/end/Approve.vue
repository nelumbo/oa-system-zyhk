<template>
    <divTable :columnObj="base.column" :tableData="base.offices" :allShow="true" />

    <el-dialog v-model="baseDLC.dialogVisible" title="审核" width="50%" :show-close="false">
        <el-row v-for="employee in baseDLC.employees">
            <el-col :span="6" :offset="3">
                员工：<el-input v-model="employee.name" readonly />
            </el-col>
            <el-col :span="6" :offset="3">
                年度提成：<el-input v-model="employee.yearMoney" readonly />
            </el-col>
        </el-row>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="baseDLC.pass" :disabled="baseDLC.submitDisabled"
                        style="margin-right: 250px;">通过</el-button>
                    <el-button type="danger" @click="baseDLC.reject" :disabled="baseDLC.submitDisabled">拒绝</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { reactive, onBeforeMount } from 'vue'
import { OfficeIsSetSubmitItems } from '@/utils/magic'
import { queryAllOffice } from "@/api/office"
import { queryAllEmployee } from "@/api/employee"
import { settApprove } from "@/api/settlement"
import { message } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    offices: [],
    column: {
        headers: [
            {
                prop: "number",
                label: "编号",
            },
            {
                prop: "name",
                label: "名称",
            },
            {
                prop: "lastYearMoney",
                label: "结算提成",
            },
            {
                type: "transform",
                prop: "isSetSubmit",
                items: OfficeIsSetSubmitItems,
                label: "状态",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.isSetSubmit == 1) {
                                return true
                            }
                            return false
                        },
                        label: "审批",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openBaseDLCDlalog(index, row)
                    }
                ]
            },
        ],
    },
    query: () => {
        queryAllOffice().then((res) => {
            if (res.status == 1) {
                base.offices = res.data
            }
        })
    },
    openBaseDLCDlalog: (index, row) => {
        baseDLC.model.id = row.id
        queryAllEmployee({ "officeID": row.id }).then((res) => {
            if (res.status == 1) {
                baseDLC.employees = res.data
            }
        })
        baseDLC.dialogVisible = true
    }
})

const baseDLC = reactive({
    dialogVisible: false,
    submitDisabled: false,
    employees: [],
    model: {
        id: null,
        isPass: null,
    },
    submit: () => {
        baseDLC.submitDisabled = true
        settApprove(baseDLC.model).then((res) => {
            if (res.status == 1) {
                message("审核成功", "success")
                base.query()
            } else {
                message("审核失败", "error")
            }
            baseDLC.dialogVisible = false
            employees: []
            baseDLC.model = {
                id: null,
                isPass: null,
            }
            baseDLC.submitDisabled = false
        })
    },
    pass: () => {
        baseDLC.model.isPass = true
        baseDLC.submit()
    },
    reject: () => {
        baseDLC.model.isPass = false
        baseDLC.submit()
    },
})

onBeforeMount(() => {
    base.query()
})

</script>