<template>
    <div>
        <el-row :gutter="10">
            <el-col :span="6" :offset="2">
                <el-select v-model="base.model.officeID" placeholder="办事处" clearable style="width: 100%;">
                    <el-option v-for="item in base.offices" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.name" placeholder="员工" clearable maxlength="25" />
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.phone" placeholder="手机号" clearable maxlength="25" />
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

        <el-dialog v-model="add.dialogVisible" title="员工添加" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="150px" :rules="rules" ref="addForm">
                <el-form-item label="办事处" prop="officeID">
                    <el-select v-model="add.model.officeID" clearable>
                        <el-option v-for="office in base.offices" :key="office.id" :label="office.name"
                            :value="office.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="姓名" prop="name">
                    <el-input v-model.trim="add.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="编号">
                    <el-input v-model.trim="add.model.number" maxlength="50" />
                </el-form-item>
                <el-form-item label="手机号" prop="phone">
                    <el-input v-model.trim="add.model.phone" maxlength="50" />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model.trim="add.model.wechatID" maxlength="50" />
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model.trim="add.model.email" maxlength="50" />
                </el-form-item>
                <el-form-item label="初始补助额度" prop="money">
                    <el-input-number v-model="add.model.money" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="每月总部补助额度" prop="credit">
                    <el-input-number v-model="add.model.credit" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="每月办事处补助额度" prop="officeCredit">
                    <el-input-number v-model="add.model.officeCredit" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="职位">
                    <el-checkbox-group v-model="add.model.roles" style="width:100%">
                        <el-row>
                            <el-col :span="6" v-for="role in add.roles" :key="role.id" style="margin-bottom: 3px;">
                                <el-checkbox :label="role">
                                    {{ role.name }}
                                </el-checkbox>
                            </el-col>
                        </el-row>
                    </el-checkbox-group>
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

        <el-dialog v-model="view.dialogVisible" title="员工查看" width="50%" :show-close="false">
            <el-form :model="view.model" label-width="150px">
                <el-form-item label="办事处">
                    <el-input v-model="view.model.office.name" disabled />
                </el-form-item>
                <el-form-item label="姓名">
                    <el-input v-model="view.model.name" disabled />
                </el-form-item>
                <el-form-item label="编号">
                    <el-input v-model="view.model.number" disabled />
                </el-form-item>
                <el-form-item label="手机号">
                    <el-input v-model="view.model.phone" disabled />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model="view.model.wechatID" disabled />
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model="view.model.email" disabled />
                </el-form-item>
                <el-form-item label="初始补助额度">
                    <el-input v-model.trim="view.model.money" disabled />
                </el-form-item>
                <el-form-item label="每月总部补助额度">
                    <el-input v-model.trim="view.model.credit" disabled />
                </el-form-item>
                <el-form-item label="每月办事处补助额度">
                    <el-input v-model.trim="view.model.officeCredit" disabled />
                </el-form-item>
                <el-form-item label="职位">
                    <el-row style="width:80%">
                        <el-col :span="6" v-for="role in view.model.roles" :key="role.id">
                            {{ role.name }}
                        </el-col>
                    </el-row>
                </el-form-item>
            </el-form>
        </el-dialog>

        <el-dialog v-model="edit.baseDialogVisible" title="员工基础信息编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="150px" :rules="rules" ref="editForm">
                <el-form-item label="姓名" prop="name">
                    <el-input v-model.trim="edit.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="手机号" prop="phone">
                    <el-input v-model.trim="edit.model.phone" maxlength="50" />
                </el-form-item>
                <el-form-item label="微信号" prop="wechatID">
                    <el-input v-model.trim="edit.model.wechatID" maxlength="50" />
                </el-form-item>
                <el-form-item label="邮箱" prop="email">
                    <el-input v-model.trim="edit.model.email" maxlength="50" />
                </el-form-item>
                <el-form-item label="办事处">
                    <el-input v-model="edit.model.office.name" disabled />
                </el-form-item>
                <el-form-item label="编号">
                    <el-input v-model="edit.model.number" disabled />
                </el-form-item>
                <el-form-item label="初始补助额度">
                    <el-input v-model.trim="edit.model.money" disabled />
                </el-form-item>
                <el-form-item label="每月总部补助额度">
                    <el-input v-model.trim="edit.model.credit" disabled />
                </el-form-item>
                <el-form-item label="每月办事处补助额度">
                    <el-input v-model.trim="edit.model.officeCredit" disabled />
                </el-form-item>
                <el-form-item label="职位">
                    <el-row style="width:80%">
                        <el-col :span="6" v-for="role in edit.model.roles" :key="role.id">
                            {{ role.name }}
                        </el-col>
                    </el-row>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="edit.submitBase"
                            :disabled="edit.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="edit.expenseDialogVisivle" title="员工财务信息编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="150px" :rules="rules" ref="editForm">
                <el-form-item label="办事处">
                    <el-input v-model="edit.model.office.name" disabled />
                </el-form-item>
                <el-form-item label="姓名">
                    <el-input v-model="edit.model.name" disabled />
                </el-form-item>
                <el-form-item label="编号">
                    <el-input v-model="edit.model.number" disabled />
                </el-form-item>
                <el-form-item label="手机号">
                    <el-input v-model="edit.model.phone" disabled />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model="edit.model.wechatID" disabled />
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model="edit.model.email" disabled />
                </el-form-item>
                <el-form-item label="初始补助额度" prop="money">
                    <el-input-number v-model="edit.model.money" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="每月总部补助额度" prop="credit">
                    <el-input-number v-model="edit.model.credit" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="每月办事处补助额度" prop="officeCredit">
                    <el-input-number v-model="edit.model.officeCredit" :controls="false" :min="-9999999" />
                </el-form-item>
                <el-form-item label="职位">
                    <el-row style="width:80%">
                        <el-col :span="6" v-for="role in edit.model.roles" :key="role.id">
                            {{ role.name }}
                        </el-col>
                    </el-row>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="edit.submitExpense"
                            :disabled="edit.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="edit.officeDialogVisivle" title="员工人事信息编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="150px" :rules="rules" ref="editForm">
                <el-form-item label="姓名">
                    <el-input v-model="edit.model.name" disabled />
                </el-form-item>
                <el-form-item label="手机号">
                    <el-input v-model="edit.model.phone" disabled />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model="edit.model.wechatID" disabled />
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model="edit.model.email" disabled />
                </el-form-item>
                <el-form-item label="初始补助额度">
                    <el-input v-model.trim="edit.model.money" disabled />
                </el-form-item>
                <el-form-item label="每月总部补助额度">
                    <el-input v-model.trim="edit.model.credit" disabled />
                </el-form-item>
                <el-form-item label="每月办事处补助额度">
                    <el-input v-model.trim="edit.model.officeCredit" disabled />
                </el-form-item>
                <el-form-item label="办事处" prop="officeID">
                    <el-select v-model="edit.model.officeID" clearable>
                        <el-option v-for="office in base.offices" :key="office.id" :label="office.name"
                            :value="office.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="编号" prop="number">
                    <el-input v-model.trim="edit.model.number" maxlength="50" />
                </el-form-item>
                <el-form-item label="职位">
                    <el-checkbox-group v-model="edit.model.roles" style="width:100%">
                        <el-row>
                            <el-col :span="6" v-for="role in edit.roles" :key="role.id">
                                <el-checkbox :label="role">
                                    {{ role.name }}
                                </el-checkbox>
                            </el-col>
                        </el-row>
                    </el-checkbox-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="edit.submitOffice"
                            :disabled="edit.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="reset.dialogVisible" title="重置密码" width="50%" :show-close="false">
            <h1>是否确定重置【{{ reset.model.name }}】的密码为【手机号】？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="reset.submit" :disabled="reset.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="del.dialogVisible" title="员工停用" width="50%" :show-close="false">
            <h1>是否确定停用【{{ del.model.name }}】？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="del.submit" :disabled="del.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { addEmployee, delEmployee, editEmployeeBase, editEmployeeExpense, editEmployeeOffice, queryEmployee, queryEmployees, resetPwd } from "@/api/employee";
import { queryAllRole } from "@/api/role";
import { queryAllOffice } from "@/api/office";
import { message, messageForCRUD } from '@/components/divMessage/index'
import { reg_money } from '@/utils/regex'

import divTable from '../../components/divTable/index.vue'

const addForm = ref(null)
const editForm = ref(null)
const rules = reactive({
    officeID: [
        { required: true, message: '请选择办事处', trigger: 'blur' },
    ],
    name: [
        { required: true, message: '请输入员工名称', trigger: 'blur' },
    ],
    phone: [
        { required: true, message: '请输入手机号', trigger: 'blur' },
    ],
    money: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    credit: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    officeCredit: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
})

const base = reactive({
    offices: [],
    model: {
        officeID: null,
        name: "",
        phone: "",
    },
    column: {
        headers: [
            {
                prop: "name",
                label: "姓名",
                width: "15%",
            },
            {
                prop: "office.name",
                label: "办事处",
                width: "15%",
            },
            {
                prop: "number",
                label: "编号",
                width: "10%",
            },
            {
                prop: "phone",
                label: "电话",
                width: "20%",
            },
            {
                type: "operation",
                label: "操作",
                width: "40%",
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
                            return true
                        },
                        label: "基础编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditBaseDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "财务编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditExpenseDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "人事编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditOfficeDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "重置密码",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openResetDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "停用",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openDelDialog(index, row)
                    }
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
        queryEmployees(base.model, base.pageData).then((res) => {
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
    openViewDialog: (index, row) => {
        queryEmployee(row.id).then((res) => {
            if (res.status == 1) {
                view.model = res.data
            }
        })
        view.dialogVisible = true
    },
    openEditBaseDialog: (index, row) => {
        queryEmployee(row.id).then((res) => {
            if (res.status == 1) {
                edit.model = res.data
            }
        })
        edit.baseDialogVisible = true
    },
    openEditExpenseDialog: (index, row) => {
        queryEmployee(row.id).then((res) => {
            if (res.status == 1) {
                edit.model = res.data
            }
        })
        edit.expenseDialogVisivle = true
    },
    openEditOfficeDialog: (index, row) => {
        queryAllRole().then((res) => {
            if (res.status == 1) {
                edit.roles = res.data
            }
        })
        queryEmployee(row.id).then((res) => {
            if (res.status == 1) {
                edit.model = res.data
                if (edit.model.officeID == 0) {
                    edit.model.officeID = null
                }
            }
        })
        edit.officeDialogVisivle = true
    },
    openResetDialog: (index, row) => {
        reset.model.id = row.id
        reset.model.name = row.name
        reset.dialogVisible = true
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
        officeID: null,
        name: "",
        phone: "",
        wechatID: "",
        email: "",
        number: "",
        money: 0,
        credit: 0,
        officeCredit: 0,
        roles: [],
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addEmployee(add.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add.model.name, "添加成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(add.model.name, "添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        officeID: null,
                        name: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                        number: "",
                        money: 0,
                        credit: 0,
                        officeCredit: 0,
                        roles: [],
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    }
})

const view = reactive({
    dialogVisible: false,
    model: {
        name: "",
        phone: "",
        wechatID: "",
        email: "",
        number: "",
        money: 0,
        credit: 0,
        officeCredit: 0,
        office: {
            name: ""
        },
        roles: [],
    },
})

const edit = reactive({
    baseDialogVisible: false,
    expenseDialogVisivle: false,
    officeDialogVisivle: false,
    submitDisabled: false,
    roles: [],
    model: {
        id: null,
        officeID: null,
        name: "",
        phone: "",
        wechatID: "",
        email: "",
        number: "",
        money: 0,
        credit: 0,
        officeCredit: 0,
        office: {
            name: ""
        },
        roles: [],
    },
    submitBase: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editEmployeeBase(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.baseDialogVisible = false
                    edit.model = {
                        id: null,
                        officeID: null,
                        name: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                        number: "",
                        money: 0,
                        credit: 0,
                        officeCredit: 0,
                        office: {
                            name: ""
                        },
                        roles: [],
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    },
    submitExpense: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editEmployeeExpense(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        // base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.expenseDialogVisivle = false
                    edit.model = {
                        id: null,
                        officeID: null,
                        name: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                        number: "",
                        money: 0,
                        credit: 0,
                        officeCredit: 0,
                        office: {
                            name: ""
                        },
                        roles: [],
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    },
    submitOffice: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editEmployeeOffice(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.officeDialogVisivle = false
                    edit.model = {
                        id: null,
                        officeID: null,
                        name: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                        number: "",
                        money: 0,
                        credit: 0,
                        officeCredit: 0,
                        office: {
                            name: ""
                        },
                        roles: [],
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const reset = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
        name: "",
    },
    submit: () => {
        reset.submitDisabled = true
        resetPwd(reset.model.id).then((res) => {
            if (res.status == 1) {
                messageForCRUD(reset.model.name, "重置密码成功", "success")
            } else {
                messageForCRUD(reset.model.name, "重置密码失败", "error")
            }
            reset.dialogVisible = false
            reset.model = {}
            reset.submitDisabled = false
        })
    }
})

const del = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
        name: "",
    },
    submit: () => {
        del.submitDisabled = true
        delEmployee(del.model.id).then((res) => {
            if (res.status == 1) {
                messageForCRUD(del.model.name, "停用成功", "success")
                base.query()
            } else {
                messageForCRUD(del.model.name, "停用失败", "error")
            }
            del.dialogVisible = false
            del.model = {
                id: 0,
                name: "",
            }
            del.submitDisabled = false
        })
    }
})

onBeforeMount(() => {
    base.query()
    queryAllOffice().then((res) => {
        if (res.status == 1) {
            base.offices = res.data
        }
    })
})
</script>