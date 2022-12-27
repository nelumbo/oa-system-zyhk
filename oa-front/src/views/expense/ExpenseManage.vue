<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-select v-model="queryObj.type" placeholder="报销类型" clearable style="width: 100%;">
                            <el-option v-for="item in expenseTypeItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-col>
                    <el-col :span="6">
                        <el-select v-model="queryObj.status" placeholder="状态" clearable style="width: 100%;">
                            <el-option v-for="item in expenseStatusItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-col>
                    <el-col :span="6">
                        <el-date-picker v-model="queryObj.startDate" type="date" placeholder="开始时间"
                            style="width: 100%;" />
                    </el-col>
                    <el-col :span="6">
                        <el-date-picker v-model="queryObj.endDate" type="date" placeholder="结束时间"
                            style="width: 100%;" />
                    </el-col>
                </el-row>
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="query">查询</el-button>
            </el-col>
            <el-col :span="1">
                <el-button type="success" @click="addDialogVisible = true">发起</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="columnObj" :tableData="tableData" :pageObj="pageObj" />

        <el-dialog v-model="addDialogVisible" title="发起报销" width="50%" :show-close="false">
            <el-form :model="form" label-width="100px">
                <el-form-item label="报销类型">
                    <el-select v-model="addObj.type" placeholder="请选择你的报销类型">
                        <el-option v-for="expenseType in expenseTypeItems" :label="expenseType.label"
                            :value="expenseType.value" />
                    </el-select>
                </el-form-item>
                <el-form-item label="金额">
                    <el-input v-model.number="addObj.amount" />
                </el-form-item>
                <el-form-item label="报销原因">
                    <el-input v-model="addObj.text" type="textarea" :autosize="{ minRows: 6, maxRows: 20 }" />
                </el-form-item>
            </el-form>
            <div style="text-align: center;">
                <el-button type="primary" @click="onSubmit">提交</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { expenseTypeItems, expenseStatusItems } from '@/utils/magic'

import divTable from '@/components/divTable/index.vue'

const addDialogVisible = ref(false)

const queryObj = reactive({
    employeeUID: "",
    type: null,
    status: null,
    startDate: "",
    endDate: "",
})

const addObj = reactive({
    type: "",
    amount: 0,
    text: "",
})

const columnObj = {
    headers: [
        {
            prop: "expenseType.text",
            label: "类型",
        },
        {
            prop: "amount",
            label: "金额(元)",
        },
        {
            prop: "text",
            label: "发起原因",
        },
        {
            prop: "CreatedAt",
            label: "发起时间",
        },
        {
            prop: "approver1.name",
            label: "审批",
        },
        {
            prop: "approver2.name",
            label: "财务",
        },
        {
            prop: "approver3.name",
            label: "出纳",
        },
        {
            prop: "UpdatedAt",
            label: "最后处理时间",
        },
        {
            prop: "status",
            label: "状态",
        },
        {
            type: "operation",
            label: "操作",
            operations: []
        },
    ],
}

const tableData = []

const pageObj = {
    total: 0,
    pageData: {
        page: 0,
        size: 10
    }
}

function query() {
    alert(JSON.stringify(queryObj))
}

</script>

<style>
.el-row {
    margin-bottom: 20px;
}
</style>