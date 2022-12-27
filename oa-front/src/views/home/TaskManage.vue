<template>
    <div>
        <el-row>
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6"></el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.no" placeholder="合同编号" />
                    </el-col>
                    <el-col :span="6">
                        <el-select v-model="queryObj.status" placeholder="任务状态" clearable style="width: 100%;">
                            <el-option v-for="item in taskStatusItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-col>
                    <el-col :span="1">
                        <el-button type="primary" @click="query">查询</el-button>
                    </el-col>
                </el-row>
            </el-col>
        </el-row>
        <divTable :columnObj="columnObj" :tableData="tableData" :pageObj="pageObj" />
    </div>
</template>

<script setup>
import { reactive } from 'vue'
import { taskStatusItems } from '../../utils/magic'

import divTable from '../../components/divTable/index.vue'

const queryObj = reactive({
    employeeUID: "",
    no: "",
    status: null,
})

const columnObj = {
    headers: [
        {
            prop: "no",
            label: "合同编号",
        },
        {
            prop: "product.name",
            label: "产品",
        },
        {
            prop: "totalPrice",
            label: "总价",
        },
        {
            prop: "number",
            label: "需求数量",
        },
        {
            prop: "product.numberCount",
            label: "库存数量",
        },
        {
            prop: "preDeposit",
            label: "单位",
        },
        {
            prop: "unit",
            label: "总金额",
        },
        {
            prop: "contract.estimatedDeliveryDate",
            label: "合同交货时间",
        },
        {
            prop: "employee",
            label: "负责人",
        },
        {
            prop: "status",
            label: "状态",
        },
        {
            prop: "startDate",
            label: "开始时间",
        },
        {
            prop: "days",
            label: "限时天数",
        },
        {
            prop: "realEndDate",
            label: "实际提交时间",
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