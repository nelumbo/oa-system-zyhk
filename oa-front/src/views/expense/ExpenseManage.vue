<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-select v-model="base.model.type" placeholder="报销类型" clearable style="width: 100%;">
                    <el-option v-for="item in expenseTypeItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                    <el-option v-for="item in expenseStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-date-picker v-model="base.model.startDate" type="date" placeholder="开始时间" style="width: 100%;" />
            </el-col>
            <el-col :span="5">
                <el-date-picker v-model="base.model.endDate" type="date" placeholder="结束时间" style="width: 100%;" />
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

        <el-dialog v-model="add.dialogVisible" title="发起报销" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="80px" :rules="rules" ref="addForm">
                <el-form-item label="报销类型" prop="type">
                    <el-select v-model="add.model.type" placeholder="请选择你的报销类型">
                        <el-option v-for="expenseType in expenseTypeItems" :label="expenseType.label"
                            :value="expenseType.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="金额" prop="amount">
                    <el-input-number v-model="add.model.amount" :controls="false" :min="0" style="width:30%" />
                </el-form-item>
                <el-form-item label="报销原因">
                    <el-input v-model="add.model.text" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
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

        <el-dialog v-model="del.dialogVisible" title="报销删除" width="50%" :show-close="false">
            <h1>是否确定删除该条申请记录？</h1>
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
import { expenseTypeItems, expenseStatusItems } from '@/utils/magic'
import { queryMyExpenses } from "@/api/my"
import { addExpense, delExpense } from "@/api/expense"
import { message } from '@/components/divMessage/index'
import { reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const rules = reactive({
    type: [
        { required: true, message: '请选择报销类型', trigger: 'blur' },
    ],
    amount: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ]
})

const base = reactive({
    model: {
        type: null,
        status: null,
        startDate: "",
        endDate: "",
    },
    column: {
        headers: [
            {
                type: "transform",
                prop: "type",
                items: expenseTypeItems,
                label: "类型",
                width: "10%",
            },
            {
                prop: "createDate",
                label: "发起时间",
                width: "10%",
            },
            {
                prop: "amount",
                label: "金额(元)",
                width: "10%",
            },
            {
                type: "textarea",
                prop: "text",
                label: "发起原因",
                width: "20%",
            },
            {
                prop: "approver.name",
                label: "审批",
                width: "8%",
            },
            {
                prop: "finance.name",
                label: "财务",
                width: "8%",
            },
            {
                prop: "cashier.name",
                label: "出纳",
                width: "8%",
            },
            {
                prop: "approveDate",
                label: "审批时间",
                width: "8%",
            },
            {
                type: "transform",
                prop: "status",
                items: expenseStatusItems,
                label: "状态",
                width: "8%",
            },
            {
                type: "operation",
                label: "操作",
                width: "10%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == -1) {
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
        queryMyExpenses(base.model, base.pageData).then((res) => {
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
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.model.name = row.name
        del.dialogVisible = true
    }
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        type: "",
        amount: 0,
        text: "",
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addExpense(add.model).then((res) => {
                    if (res.status == 1) {
                        message("发起成功", "success")
                        base.query()
                    } else {
                        message("发起失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        type: "",
                        amount: 0,
                        text: "",
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

const del = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
    },
    submit: () => {
        del.submitDisabled = true
        delExpense(del.model.id).then((res) => {
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