<template>
    <el-row :gutter="20">
        <el-col :span="6" :offset="8">
            <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                <el-option v-for="item in bidBondStatusItems" :key="item.value" :label="item.label"
                    :value="item.value" />
            </el-select>
        </el-col>
        <el-col :span="1">
            <el-button type="primary" @click="base.query">查询</el-button>
        </el-col>
        <el-col :span="1">
            <el-button type="success" @click="base.openAddDialog">发起</el-button>
        </el-col>
    </el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

    <el-dialog v-model="add.dialogVisible" title="投标保证金录入" width="50%" :show-close="false">
        <el-form :model="add.model" label-width="60px" :rules="rules" ref="addForm">
            <el-form-item label="金额" prop="money">
                <el-input-number v-model="add.model.money" :controls="false" :min="0" />
            </el-form-item>
            <el-form-item label="备注" prop="createRemark">
                <el-input v-model="add.model.createRemark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
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

    <el-dialog v-model="edit.dialogVisible" title="投标保证金编辑" width="50%" :show-close="false">
        <el-form :model="edit.model" label-width="60px" :rules="rules" ref="editForm">
            <el-form-item label="金额" prop="money">
                <el-input-number v-model="edit.model.money" :controls="false" :min="0" />
            </el-form-item>
            <el-form-item label="备注" prop="createRemark">
                <el-input v-model="edit.model.createRemark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
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

    <el-dialog v-model="del.dialogVisible" title="投标保证金删除" width="50%" :show-close="false">
        <h1>是否确定删除该条投标保证金记录？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="del.submit" :disabled="del.submitDisabled">确定</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { bidBondStatusItems } from '@/utils/magic'
import { queryMyBidbonds } from "@/api/my"
import { addBidbond, delBidbond, editBidbond } from "@/api/bidbond"
import { message } from '@/components/divMessage/index'
import { reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const editForm = ref(null)
const rules = reactive({
    money: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    createRemark: [
        { required: true, message: '请填写备注，不超过300个字！', trigger: 'blur' },
    ],
})

const base = reactive({
    model: {
        status: null,
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "录入日期",
                width: "10%",
            },
            {
                prop: "employee.name",
                label: "业务员",
                width: "8%",
            },
            {
                prop: "money",
                label: "金额",
                width: "7%",
            },
            {
                type: "textarea",
                prop: "createRemark",
                label: "业务员备注",
                width: "15%",
            },
            {
                prop: "auditor.name",
                label: "付款财务",
                width: "5%",
            },
            {
                prop: "auditDate",
                label: "付款日期",
                width: "10%",
            },
            {
                prop: "finalce.name",
                label: "回款财务",
                width: "5%",
            },
            {
                prop: "finalDate",
                label: "回款日期",
                width: "10%",
            },
            {
                type: "textarea",
                prop: "finalRemark",
                label: "回款备注",
                width: "15%",
            },
            {
                type: "transform",
                prop: "status",
                items: bidBondStatusItems,
                label: "状态",
                width: "5%",
            },
            {
                type: "operation",
                label: "操作",
                width: "10%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == -1 || row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "删除",
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
        queryMyBidbonds(base.model, base.pageData).then((res) => {
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
        add.dialogVisible = true
    },
    openEditDialog: (index, row) => {
        edit.model.id = row.id
        edit.model.money = row.money
        edit.model.createRemark = row.createRemark
        edit.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.dialogVisible = true
    }
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        money: 0,
        createRemark: "",
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addBidbond(add.model).then((res) => {
                    if (res.status == 1) {
                        message("录入成功", "success")
                        base.query()
                    } else {
                        message("录入失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        money: 0,
                        createRemark: "",
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

const edit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        money: 0,
        createRemark: "",
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editBidbond(edit.model).then((res) => {
                    if (res.status == 1) {
                        message("编辑成功", "success")
                        base.query()
                    } else {
                        message("编辑失败", "error")
                    }
                    edit.dialogVisible = false
                    edit.model = {
                        id: null,
                        money: 0,
                        createRemark: "",
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const del = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
    },
    submit: () => {
        del.submitDisabled = true
        delBidbond(del.model.id).then((res) => {
            if (res.status == 1) {
                message("删除成功", "success")
                base.query()
            } else {
                message("删除失败", "error")
            }
            del.dialogVisible = false
            del.model = {
                id: 0,
            }
            del.submitDisabled = false
        })
    }
})

onBeforeMount(() => {
    base.query()
})
</script>