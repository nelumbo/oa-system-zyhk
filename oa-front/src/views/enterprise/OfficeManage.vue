<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="6" :offset="8">
                <el-input v-model="base.model.name" placeholder="名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
            <el-col :span="1">
                <el-button type="success" @click="base.openAddDialog">添加</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="add.dialogVisible" title="办事处/部门添加" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="120px" :rules="rules" ref="addForm">
                <el-form-item label="编号">
                    <el-input v-model.trim="add.model.number" maxlength="50" />
                </el-form-item>
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="add.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="业务费金额" prop="businessMoney">
                    <el-input-number v-model="add.model.businessMoney" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="提成金额" prop="money">
                    <el-input-number v-model="add.model.money" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="年底提成金额" prop="moneyCold">
                    <el-input-number v-model="add.model.moneyCold" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="今年目标金额" prop="taskLoad">
                    <el-input-number v-model="add.model.taskLoad" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="今年完成金额" prop="targetLoad">
                    <el-input-number v-model="add.model.targetLoad" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="默认职位">
                    <el-select v-model="add.model.roleID" clearable>
                        <el-option v-for="role in add.roles" :key="role.id" :label="role.name" :value="role.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="列入排行榜">
                    <el-radio-group v-model="add.model.isRanking">
                        <el-radio :label="true">是</el-radio>
                        <el-radio :label="false">否</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="add.submit" :disabled="add.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="edit.dialogVisible" title="办事处/部门编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="120px" :rules="rules" ref="editForm">
                <el-form-item label="编号">
                    <el-input v-model.trim="edit.model.number" maxlength="50" />
                </el-form-item>
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="edit.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="业务费金额" prop="businessMoney">
                    <el-input-number v-model="edit.model.businessMoney" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="提成金额" prop="money">
                    <el-input-number v-model="edit.model.money" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="年底提成金额" prop="moneyCold">
                    <el-input-number v-model="edit.model.moneyCold" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="今年目标金额" prop="taskLoad">
                    <el-input-number v-model="edit.model.taskLoad" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="今年完成金额" prop="targetLoad">
                    <el-input-number v-model="edit.model.targetLoad" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="默认职位">
                    <el-select v-model="edit.model.roleID" clearable>
                        <el-option v-for="role in edit.roles" :key="role.id" :label="role.name" :value="role.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="列入排行榜">
                    <el-radio-group v-model="edit.model.isRanking">
                        <el-radio :label="true">是</el-radio>
                        <el-radio :label="false">否</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="edit.submit" :disabled="edit.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { addOffice, editOffice, queryOffices } from "@/api/office"
import { queryAllRole } from '@/api/role'
import { message, messageForCRUD } from '@/components/divMessage/index'
import { reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const editForm = ref(null)
const rules = reactive({
    name: [
        { required: true, message: '请输入办事处名称', trigger: 'blur' },
    ],
    businessMoney: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    money: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    moneyCold: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    taskLoad: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    targetLoad: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
})

const base = reactive({
    model: {
        name: "",
    },
    column: {
        headers: [
            {
                prop: "number",
                label: "办事处编号",
            },
            {
                prop: "name",
                label: "办事处名称",
            },
            {
                prop: "businessMoney",
                label: "可用业务费用金额",
            },
            {
                prop: "money",
                label: "可提成金额",
            },
            {
                prop: "moneyCold",
                label: "年底提成金额",
            },
            {
                prop: "taskLoad",
                label: "今年目标金额",
            },
            {
                prop: "targetLoad",
                label: "今年完成金额",
            },
            {
                prop: "role.name",
                label: "默认职位",
            },
            {
                type: "boolean",
                prop: "isRanking",
                label: "入榜",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditDialog(index, row)
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
        queryOffices(base.model, base.pageData).then((res) => {
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
    openAddDialog: () => {
        queryAllRole().then((res) => {
            if (res.status == 1) {
                add.roles = res.data
            }
        })
        add.dialogVisible = true
    },
    openEditDialog: (index, row) => {
        queryAllRole().then((res) => {
            if (res.status == 1) {
                edit.roles = res.data
            }
        })
        edit.model = row
        if (edit.model.roleID == 0) {
            edit.model.roleID = null
        }
        edit.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.model.name = row.name
        del.dialogVisible = true
    }
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    roles: [],
    model: {
        number: "",
        name: "",
        businessMoney: 0,
        money: 0,
        moneyCold: 0,
        taskLoad: 0,
        targetLoad: 0,
        roleID: null,
        isRanking: true,
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addOffice(add.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add.model.name, "添加成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(add.model.name, "添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        number: "",
                        name: "",
                        businessMoney: 0,
                        money: 0,
                        moneyCold: 0,
                        taskLoad: 0,
                        targetLoad: 0,
                        roleID: null,
                        isRanking: true,
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const edit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    roles: [],
    model: {
        number: "",
        name: "",
        businessMoney: 0,
        money: 0,
        moneyCold: 0,
        taskLoad: 0,
        targetLoad: 0,
        roleID: null,
        isRanking: true,
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editOffice(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.dialogVisible = false
                    edit.model = {
                        number: "",
                        name: "",
                        businessMoney: 0,
                        money: 0,
                        moneyCold: 0,
                        taskLoad: 0,
                        targetLoad: 0,
                        roleID: null,
                        isRanking: true,
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

onBeforeMount(() => {
    base.query()
})
</script>