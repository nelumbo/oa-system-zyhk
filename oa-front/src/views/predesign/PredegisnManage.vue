<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="6" :offset="5">
                <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                    <el-option v-for="item in predesignStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.employee.name" placeholder="业务员" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="view.dialogVisible" title="查看" width="75%" :show-close="false">
            <el-form :model="view.model" label-width="100px">
                <el-form-item label="发起日期">
                    <el-input v-model.trim="view.model.createDate" readonly />
                </el-form-item>
                <el-form-item label="办事处">
                    <el-input v-model.trim="view.model.employee.office.name" readonly />
                </el-form-item>
                <el-form-item label="业务员">
                    <el-input v-model.trim="view.model.employee.name" readonly />
                </el-form-item>
                <el-form-item label="业务员备注">
                    <el-input v-model.trim="view.model.remark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="审核">
                    <el-input v-model.trim="view.model.auditor.name" readonly />
                </el-form-item>
                <el-form-item label="审核日期">
                    <el-input v-model.trim="view.model.auditDate" readonly />
                </el-form-item>
                <el-form-item label="完成日期">
                    <el-input v-model.trim="view.model.finalDate" readonly />
                </el-form-item>
                <el-form-item label="状态">
                    <el-input v-model.trim="view.statusString" readonly />
                </el-form-item>
            </el-form>
            <el-divider></el-divider>
            <divTable :columnObj="view.column" :tableData="view.model.predesignTasks" :allShow="true" />
        </el-dialog>

        <el-dialog v-model="approve.dialogVisible" title="审核" width="50%" :show-close="false">
            <el-form :model="approve.model" label-width="100px" :rules="rules" ref="approveForm">
                <el-form-item label="发起日期">
                    <el-input v-model.trim="approve.model.createDate" disabled />
                </el-form-item>
                <el-form-item label="办事处">
                    <el-input v-model.trim="approve.model.employee.office.name" disabled />
                </el-form-item>
                <el-form-item label="业务员">
                    <el-input v-model.trim="approve.model.employee.name" disabled />
                </el-form-item>
                <el-form-item label="业务员备注">
                    <el-input v-model.trim="approve.model.remark" type="textarea" autosize disabled />
                </el-form-item>
                <el-divider></el-divider>
                <el-form-item label="办事处">
                    <el-select v-model="approve.officeID" @change="approve.queryEmployees">
                        <el-option v-for="office in approve.offices" :key="office.id" :label="office.name"
                            :value="office.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="员工" prop="predesignTask.employeeID">
                    <el-select v-model="approve.model.predesignTask.employeeID" clearable>
                        <el-option v-for="employee in approve.employees" :key="employee.id" :label="employee.name"
                            :value="employee.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="工作天数" prop="predesignTask.days">
                    <el-input-number v-model="approve.model.predesignTask.days" :controls="false" :min="0" :max="999"
                        style="width:30%" />
                </el-form-item>
                <el-form-item label="设计需求">
                    <el-input v-model="approve.model.predesignTask.createRemark" type="textarea"
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
import { user } from '@/pinia/modules/user'
import { computed, ref, reactive, onBeforeMount } from 'vue'
import { predesignStatusItems, predesignTaskStatusItems } from '@/utils/magic'
import { queryAllOffice } from "@/api/office"
import { queryAllEmployee } from "@/api/employee"
import { approvePredesign, queryPredesign, queryPredesigns } from "@/api/predesign"
import { message } from '@/components/divMessage/index'
import { reg_number } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const approveForm = ref(null)
const rules = reactive({
    predesignTask: {
        employeeID: [
            { required: true, message: '请选择员工', trigger: 'blur' },
        ],
        days: [
            { required: true, pattern: reg_number, message: '请输入有效数字', trigger: 'blur' }
        ],
    }
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
                label: "发起日期",
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
                prop: "auditor.name",
                label: "审核人",
            },
            {
                prop: "auditDate",
                label: "审核日期",
            },
            {
                type: "transform",
                prop: "status",
                items: predesignStatusItems,
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
                            if (user().my.pids.includes('22') && row.status == 1) {
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
        queryPredesigns(base.model, base.pageData).then((res) => {
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
        queryPredesign(row.id).then((res) => {
            if (res.status == 1) {
                view.model = res.data
            }
        })
        view.dialogVisible = true
    },
    openApproveDialog: (index, row) => {
        queryAllOffice().then((res) => {
            if (res.status == 1) {
                approve.offices = res.data
            }
        })
        approve.model = row
        approve.model.predesignTask.employeeID = null
        approve.model.predesignTask.createRemark = row.remark
        approve.dialogVisible = true
    }
})

const view = reactive({
    dialogVisible: false,
    model: {
        createDate: "",
        remark: "",
        auditDate: "",
        finalDate: "",
        status: "",
        predesignTasks: [],
        employee: {
            name: "",
            office: {
                name: "",
            }
        },
        auditor: {
            name: "",
        },
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
        ]
    },
    statusString: computed(() => {
        var temp = "";
        predesignStatusItems.some((item) => {
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
    offices: [],
    employees: [],
    officeID: null,
    model: {
        id: null,
        createDate: "",
        remark: "",
        employee: {
            name: "",
            office: {
                name: "",
            }
        },
        predesignTask: {
            employeeID: null,
            createRemark: "",
            days: null,
        },
        isPass: null,
    },
    queryEmployees: () => {
        approve.model.predesignTask.employeeID = null,
            queryAllEmployee({ "officeID": approve.officeID }).then((res => {
                if (res.status == 1) {
                    approve.employees = res.data
                }
            }))
    },
    submit: () => {
        approveForm.value.validate((valid) => {
            if (valid) {
                approve.submitDisabled = true
                approvePredesign(
                    {
                        "id": approve.model.id,
                        "remark": approve.model.remark,
                        "isPass": approve.model.isPass,
                        "predesignTask": {
                            "employeeID": approve.model.predesignTask.employeeID,
                            "days": approve.model.predesignTask.days,
                            "createRemark": approve.model.predesignTask.createRemark,
                        }
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
                        remark: "",
                        employee: {
                            name: "",
                            office: {
                                name: "",
                            }
                        },
                        predesignTask: {
                            employeeID: null,
                            createRemark: "",
                            days: null,
                        },
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