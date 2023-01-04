<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="6" :offset="2">
                <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                    <el-option v-for="item in base.offices" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-select v-model="base.model.employee.officeID" placeholder="办事处" clearable style="width: 100%;">
                    <el-option v-for="item in base.offices" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.employee.name" placeholder="业务员" clearable :maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="approve.dialogVisible" title="审核" width="50%" :show-close="false">
            <el-form :model="approve.model" label-width="100px">
                <el-form-item label="办事处">
                    <el-input v-model.trim="approve.model.employee.office.name" readonly />
                </el-form-item>
                <el-form-item label="业务员">
                    <el-input v-model.trim="approve.model.employee.name" readonly />
                </el-form-item>
                <el-form-item label="金额">
                    <el-input v-model.trim="approve.model.money" readonly />
                </el-form-item>
                <el-form-item label="业务员备注">
                    <el-input v-model.trim="approve.model.createRemark" type="textarea" autosize readonly />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="approve.pass" :disabled="approve.submitDisabled"
                            style="margin-right: 250px;">通过</el-button>
                        <el-button type="danger" @click="approve.reject"
                            :disabled="approve.submitDisabled">拒绝</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="final.dialogVisible" title="回款" width="50%" :show-close="false">
            <el-form :model="final.model" label-width="100px">
                <el-form-item label="办事处">
                    <el-input v-model.trim="final.model.employee.office.name" disabled />
                </el-form-item>
                <el-form-item label="业务员">
                    <el-input v-model.trim="final.model.employee.name" disabled />
                </el-form-item>
                <el-form-item label="金额">
                    <el-input v-model.trim="final.model.money" disabled />
                </el-form-item>
                <el-form-item label="业务员备注">
                    <el-input v-model.trim="final.model.createRemark" type="textarea" autosize disabled />
                </el-form-item>
                <el-divider />
                <el-form-item label="回款备注">
                    <el-input v-model="final.model.finalRemark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="final.pass" :disabled="final.submitDisabled"
                            style="margin-right: 250px;">通过</el-button>
                        <el-button type="danger" @click="final.reject" :disabled="final.submitDisabled">拒绝</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { reactive, onBeforeMount } from 'vue'
import { bidBondStatusItems } from '@/utils/magic'
import { queryAllOffice } from "@/api/office"
import { approveBidbond, queryBidbonds } from "@/api/bidbond"
import { message } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    offices: [],
    model: {
        status: null,
        employee: {
            officeID: null,
            name: "",
        },
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "录入时间",
            },
            {
                prop: "employee.office.name",
                label: "办事处",
            },
            {
                prop: "employee.name",
                label: "业务员",
            },
            {
                prop: "money",
                label: "金额",
            },
            {
                type: "textarea",
                prop: "createRemark",
                label: "业务员备注",
            },
            {
                prop: "auditor.name",
                label: "审核",
            },
            {
                prop: "auditDate",
                label: "审核日期",
            },
            {
                prop: "finalce.name",
                label: "财务",
            },
            {
                prop: "finalDate",
                label: "回款日期",
            },
            {
                type: "textarea",
                prop: "finalRemark",
                label: "回款备注",
            },
            {
                type: "transform",
                prop: "status",
                items: bidBondStatusItems,
                label: "状态",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "审核",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openApproveDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 2) {
                                return true
                            }
                            return false
                        },
                        label: "回款",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openFinalDialog(index, row)
                    },
                ]
            },
        ],
    },
    tableData: [],
    pageData: {
        total: 0,
        pageSize: 10,
        pageNo: 1
    },
    query: () => {
        queryBidbonds(base.model, base.pageData).then((res) => {
            if (res.status == 1) {
                base.tableData = res.data.data
                base.pageData.total = res.data.total
                base.pageData.pageSize = res.data.pageSize
                base.pageData.pageNo = res.data.pageNo
            } else {
                message("查询失败", "error")
            }
        })
    },
    handleSizeChange: (e) => {
        base.pageData.pageSize = e
        base.pageData.pageNo = 1
        base.query()
    },
    handleCurrentChange: (e) => {
        base.pageData.pageNo = e
        base.query()
    },
    openApproveDialog: (index, row) => {
        approve.model = row
        approve.dialogVisible = true
    },
    openFinalDialog: (index, row) => {
        final.model = row
        final.dialogVisible = true
    }
})

const approve = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        money: 0,
        createRemark: "",
        employee: {
            name: "",
            office: {
                name: "",
            }
        },
        status: null,
        isPass: null,
    },
    submit: () => {
        approve.submitDisabled = true
        approveBidbond(
            {
                "id": approve.model.id,
                "status": approve.model.status,
                "isPass": approve.model.isPass
            }
        ).then((res) => {
            if (res.status == 1) {
                message("审核成功", "success")
                base.query()
            } else {
                message("审核失败", "error")
            }
            approve.dialogVisible = false
            approve.model = {
                iid: null,
                money: 0,
                createRemark: "",
                employee: {
                    name: "",
                    office: {
                        name: "",
                    }
                },
                status: null,
                isPass: null,
            }
            approve.submitDisabled = false
        })
    },
    pass: () => {
        approve.model.isPass = true
        approve.submit()
    },
    reject: () => {
        approve.model.isPass = false
        approve.submit()
    },
})

const final = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        money: 0,
        createRemark: "",
        finalRemark: "",
        employee: {
            name: "",
            office: {
                name: "",
            }
        },
        status: null,
        isPass: null,
    },
    submit: () => {
        final.submitDisabled = true
        approveBidbond(
            {
                "id": final.model.id,
                "status": final.model.status,
                "finalRemark": final.model.finalRemark,
                "isPass": final.model.isPass
            }
        ).then((res) => {
            if (res.status == 1) {
                message("回款成功", "success")
                base.query()
            } else {
                message("回款失败", "error")
            }
            final.dialogVisible = false
            final.model = {
                id: null,
                money: 0,
                createRemark: "",
                finalRemark: "",
                employee: {
                    name: "",
                    office: {
                        name: "",
                    }
                },
                status: null,
                isPass: null,
            }
            final.submitDisabled = false
        })
    },
    pass: () => {
        final.model.isPass = true
        final.submit()
    },
    reject: () => {
        final.model.isPass = false
        final.submit()
    },
})

onBeforeMount(() => {
    queryAllOffice().then((res) => {
        if (res.status == 1) {
            base.offices = res.data
        }
    })
    base.query()
})
</script>