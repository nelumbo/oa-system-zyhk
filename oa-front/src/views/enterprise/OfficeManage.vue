<template>
    <el-row :gutter="20">
        <el-col :span="6" :offset="8">
            <el-input v-model="base.model.name" placeholder="名称" clearable maxlength="25" />
        </el-col>
        <el-col :span="1">
            <el-button type="primary" @click="base.query">查询</el-button>
        </el-col>
        <el-col :span="1" v-if="user().my.pids.includes('52')">
            <el-button type="success" @click="base.openAddDialog">添加</el-button>
        </el-col>
    </el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

    <el-dialog v-model="add.dialogVisible" title="添加" width="50%" :show-close="false">
        <el-form :model="add.model" label-width="120px" :rules="rules" ref="addForm">
            <el-form-item label="编号">
                <el-input v-model.trim="add.model.number" maxlength="50" />
            </el-form-item>
            <el-form-item label="名称" prop="name">
                <el-input v-model.trim="add.model.name" maxlength="50" />
            </el-form-item>
            <el-form-item label="今年目标金额" prop="taskLoad">
                <el-input-number v-model="add.model.taskLoad" :controls="false" :min="-9999999" />
            </el-form-item>
            <el-form-item label="默认职位">
                <el-select v-model="add.model.roleID" clearable>
                    <el-option v-for="role in add.roles" :key="role.id" :label="role.name" :value="role.id" />
                </el-select>
            </el-form-item>
            <el-form-item label="排行榜" prop="officeID">
                <el-select v-model="add.model.rankingNo" clearable>
                    <el-option v-for="officeRankingNo in OfficeRankingNoItems" :key="officeRankingNo.value"
                        :label="officeRankingNo.label" :value="officeRankingNo.value" />
                </el-select>
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

    <el-dialog v-model="editBase.dialogVisible" title="基础编辑" width="50%" :show-close="false">
        <el-form :model="editBase.model" label-width="120px" :rules="rules" ref="editBaseForm">
            <el-form-item label="编号">
                <el-input v-model.trim="editBase.model.number" maxlength="50" />
            </el-form-item>
            <el-form-item label="名称" prop="name">
                <el-input v-model.trim="editBase.model.name" maxlength="50" />
            </el-form-item>
            <el-form-item label="今年目标金额" prop="taskLoad">
                <el-input-number v-model="editBase.model.taskLoad" :controls="false" :min="-9999999" />
            </el-form-item>
            <el-form-item label="默认职位">
                <el-select v-model="editBase.model.roleID" clearable>
                    <el-option v-for="role in editBase.roles" :key="role.id" :label="role.name" :value="role.id" />
                </el-select>
            </el-form-item>
            <el-form-item label="排行榜" prop="officeID">
                <el-select v-model="editBase.model.rankingNo" clearable>
                    <el-option v-for="officeRankingNo in OfficeRankingNoItems" :key="officeRankingNo.value"
                        :label="officeRankingNo.label" :value="officeRankingNo.value" />
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="editBase.submit"
                        :disabled="editBase.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="editMoney.dialogVisible" title="费用编辑" width="50%" :show-close="false">
        <el-form :model="editMoney.model" label-width="120px" :rules="rules" ref="editMoneyForm">
            <el-form-item label="业务费金额" prop="businessMoney">
                <el-input-number v-model="editMoney.model.businessMoney" :controls="false" :min="-9999999" />
            </el-form-item>
            <el-form-item label="提成金额" prop="money">
                <el-input-number v-model="editMoney.model.money" :controls="false" :min="-9999999" />
            </el-form-item>
            <el-form-item label="年底提成金额" prop="moneyCold">
                <el-input-number v-model="editMoney.model.moneyCold" :controls="false" :min="-9999999" />
            </el-form-item>
            <el-form-item label="今年完成金额" prop="targetLoad">
                <el-input-number v-model="editMoney.model.targetLoad" :controls="false" :min="-9999999" />
            </el-form-item>
            <el-form-item label="备注" prop="remark">
                <el-input v-model="editMoney.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="100" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="editMoney.submit"
                        :disabled="editMoney.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { ref, reactive, onBeforeMount } from 'vue'
import { OfficeRankingNoItems } from '@/utils/magic'
import { addOffice, editOfficeBase, editOfficeMoney, queryOffices } from "@/api/office"
import { queryAllRole } from '@/api/role'
import { message } from '@/components/divMessage/index'
import { reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const editBaseForm = ref(null)
const editMoneyForm = ref(null)
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
    remark: [
        { required: true, message: '请填写', trigger: 'blur' },
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
                label: "编号",
                width: "7%"
            },
            {
                prop: "name",
                label: "名称",
                width: "10%"
            },
            {
                prop: "businessMoney",
                label: "业务费",
                width: "8%"
            },
            {
                prop: "money",
                label: "可提成金额",
                width: "8%"
            },
            {
                prop: "moneyCold",
                label: "年底提成金额",
                width: "8%"
            },
            {
                prop: "taskLoad",
                label: "今年目标金额",
                width: "8%"
            },
            {
                prop: "targetLoad",
                label: "今年完成金额",
                width: "8%"
            },
            {
                prop: "lastYearMoney",
                label: "结算提成",
                width: "10%"
            },
            {
                prop: "role.name",
                label: "默认职位",
                width: "8%"
            },
            {
                type: "operation",
                label: "操作",
                width: "25%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (user().my.pids.includes('53')) {
                                return true
                            }
                            return false
                        },
                        label: "基础编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditBaseDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (user().my.pids.includes('53')) {
                                return true
                            }
                            return false
                        },
                        label: "费用编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditMoneyDialog(index, row)
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
    openEditBaseDialog: (index, row) => {
        queryAllRole().then((res) => {
            if (res.status == 1) {
                editBase.roles = res.data
            }
        })

        editBase.model.id = row.id
        editBase.model.number = row.number
        editBase.model.name = row.name
        editBase.model.taskLoad = row.taskLoad
        if (row.roleID == 0) {
            editBase.model.roleID = null
        } else {
            editBase.model.roleID = row.roleID
        }
        editBase.model.rankingNo = row.rankingNo

        editBase.dialogVisible = true
    },
    openEditMoneyDialog: (index, row) => {
        editMoney.model.id = row.id
        editMoney.dialogVisible = true
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
        rankingNo: 0,
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addOffice(add.model).then((res) => {
                    if (res.status == 1) {
                        message("添加成功", "success")
                        base.query()
                    } else {
                        message("添加失败", "error")
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
                        rankingNo: 0,
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const editBase = reactive({
    dialogVisible: false,
    submitDisabled: false,
    roles: [],
    model: {
        id: null,
        number: "",
        name: "",
        taskLoad: 0,
        roleID: null,
        rankingNo: 0,
    },
    submit: () => {
        editBaseForm.value.validate((valid) => {
            if (valid) {
                editBase.submitDisabled = true
                editOfficeBase(
                    {
                        "id": editBase.model.id,
                        "name": editBase.model.name,
                        "number": editBase.model.number,
                        "taskLoad": editBase.model.taskLoad,
                        "roleID": editBase.model.roleID,
                        "rankingNo": editBase.model.rankingNo,
                    }
                ).then((res) => {
                    if (res.status == 1) {
                        message("编辑成功", "success")
                        base.query()
                    } else {
                        message("编辑失败", "error")
                    }
                    editBase.dialogVisible = false
                    editBase.model = {
                        id: null,
                        number: "",
                        name: "",
                        taskLoad: 0,
                        roleID: null,
                        rankingNo: 0,
                    }
                    editBase.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const editMoney = reactive({
    dialogVisible: false,
    submitDisabled: false,
    roles: [],
    model: {
        id: null,
        businessMoney: 0,
        money: 0,
        moneyCold: 0,
        targetLoad: 0,
        remark: "",
    },
    submit: () => {
        editMoneyForm.value.validate((valid) => {
            if (valid) {
                editMoney.submitDisabled = true
                editOfficeMoney(
                    {
                        "id": editMoney.model.id,
                        "businessMoney": editMoney.model.businessMoney,
                        "money": editMoney.model.money,
                        "moneyCold": editMoney.model.moneyCold,
                        "targetLoad": editMoney.model.targetLoad,
                        "remark": editMoney.model.remark,
                    }
                ).then((res) => {
                    if (res.status == 1) {
                        message("编辑成功", "success")
                        base.query()
                    } else {
                        message("编辑失败", "error")
                    }
                    editMoney.dialogVisible = false
                    editMoney.model = {
                        id: null,
                        businessMoney: 0,
                        money: 0,
                        moneyCold: 0,
                        targetLoad: 0,
                        remark: "",
                    }
                    editMoney.submitDisabled = false
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