<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="6" :offset="5">
                <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                    <el-option v-for="item in predesignTaskStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.employee.name" placeholder="技术员" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="view.dialogVisible" title="查看" width="50%" :show-close="false">
            <el-form :model="view.model" label-width="100px">
                <el-form-item label="开始日期">
                    <el-input v-model.trim="view.model.createDate" readonly />
                </el-form-item>
                <el-form-item label="工作天数">
                    <el-input v-model.trim="view.model.days" readonly />
                </el-form-item>
                <el-form-item label="预期完成日期">
                    <el-input v-model.trim="view.model.endDate" readonly />
                </el-form-item>
                <el-form-item label="设计需求">
                    <el-input v-model.trim="view.model.createRemark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="技术员">
                    <el-input v-model.trim="view.model.employee.name" readonly />
                </el-form-item>
                <el-form-item label="技术员备注">
                    <el-input v-model.trim="view.model.finalRemark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="提交日期">
                    <el-input v-model.trim="view.model.finalDate" readonly />
                </el-form-item>
                <el-form-item label="审核">
                    <el-input v-model.trim="view.model.auditor.name" readonly />
                </el-form-item>
                <el-form-item label="审核日期">
                    <el-input v-model.trim="view.model.approveDate" readonly />
                </el-form-item>
                <el-form-item label="状态">
                    <el-input v-model.trim="view.statusString" readonly />
                </el-form-item>
            </el-form>
        </el-dialog>

        <el-dialog v-model="approve.dialogVisible" title="审核" width="50%" :show-close="false">
            <el-form :model="approve.model" label-width="100px" :rules="rules" ref="approveForm">
                <el-form-item label="开始日期">
                    <el-input v-model.trim="approve.model.createDate" disabled />
                </el-form-item>
                <el-form-item label="工作天数">
                    <el-input v-model.trim="approve.model.days" disabled />
                </el-form-item>
                <el-form-item label="预期完成日期">
                    <el-input v-model.trim="approve.model.endDate" disabled />
                </el-form-item>
                <el-form-item label="设计需求">
                    <el-input v-model.trim="approve.model.createRemark" type="textarea" autosize disabled />
                </el-form-item>
                <el-form-item label="技术员">
                    <el-input v-model.trim="approve.model.employee.name" disabled />
                </el-form-item>
                <el-form-item label="技术员备注">
                    <el-input v-model.trim="approve.model.finalRemark" type="textarea" autosize disabled />
                </el-form-item>
                <el-form-item label="提交日期">
                    <el-input v-model.trim="approve.model.finalDate" disabled />
                </el-form-item>
                <el-divider></el-divider>
                <el-form-item label="新工作天数" prop="days">
                    <el-input-number v-model="approve.model.newDays" :controls="false" :min="0" :max="999"
                        style="width:30%" />
                </el-form-item>
                <el-form-item label="新设计需求">
                    <el-input v-model="approve.model.newCreateRemark" type="textarea"
                        :autosize="{ minRows: 3, maxRows: 9 }" maxlength="300" />
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

    </div>
</template>

<script setup>
import { computed, ref, reactive, onBeforeMount } from 'vue'
import { predesignTaskStatusItems } from '@/utils/magic'
import { approvePredesignTask, queryPredesignTasks } from "@/api/predesign_task"
import { message } from '@/components/divMessage/index'
import { reg_number } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const approveForm = ref(null)
const rules = reactive({
    days: [
        { required: true, pattern: reg_number, message: '请输入有效数字', trigger: 'blur' }
    ],
})

const base = reactive({
    model: {
        status: null,
        employee: {
            name: "",
        }
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "开始日期",
            },
            {
                prop: "endDate",
                label: "预期完成时间",
            },
            {
                prop: "finalDate",
                label: "提交日期",
            },
            {
                prop: "employee.name",
                label: "技术员",
            },
            {
                prop: "auditor.name",
                label: "审核",
            },
            {
                type: "transform",
                prop: "status",
                items: predesignTaskStatusItems,
                label: "状态",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "查看",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openViewDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 2) {
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
        queryPredesignTasks(base.model, base.pageData).then((res) => {
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
    openViewDialog: (index, row) => {
        view.model = row
        view.dialogVisible = true
    },
    openApproveDialog: (index, row) => {
        approve.model = row
        approve.model.newCreateRemark = row.createRemark
        approve.dialogVisible = true
    }
})

const view = reactive({
    dialogVisible: false,
    model: {
        createDate: "",
        days: "",
        endDate: "",
        remark: "",
        finalRemark: "",
        finalDate: "",
        approveDate: "",
        status: null,
        employee: {
            name: "",
        },
        auditor: {
            name: "",
        }
    },
    statusString: computed(() => {
        var temp = "";
        predesignTaskStatusItems.some((item) => {
            if (item.value == view.model.status) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
})

const approve = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        createDate: "",
        days: "",
        endDate: "",
        createRemark: "",
        finalRemark: "",
        finalDate: "",
        employee: {
            name: "",
        },
        newDays: 0,
        newCreateRemark: "",
        isPass: null,
    },
    submit: () => {
        approveForm.value.validate((valid) => {
            if (valid) {
                approve.submitDisabled = true
                approvePredesignTask(
                    {
                        "id": approve.model.id,
                        "newDays": approve.model.newDays,
                        "newCreateRemark": approve.model.newCreateRemark,
                        "isPass": approve.model.isPass,
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
                        id: null,
                        createDate: "",
                        days: "",
                        endDate: "",
                        createRemark: "",
                        finalRemark: "",
                        finalDate: "",
                        employee: {
                            name: "",
                        },
                        newDays: 0,
                        newCreateRemark: "",
                        isPass: null,
                    }
                    approve.submitDisabled = false
                })
            } else {
                return false;
            }
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

onBeforeMount(() => {
    base.query()
})
</script>