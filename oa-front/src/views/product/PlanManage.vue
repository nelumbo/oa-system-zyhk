<template>
    <div>
        <el-row :gutter="20" align="middle">
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-input v-model="queryObj.no" placeholder="合同编号" />
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.type" placeholder="类别" />
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.product" placeholder="物品名称" />
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.specification" placeholder="规格型号" />
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-input v-model="queryObj.buyNumber" placeholder="采购数量" />
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.unit" placeholder="单位" />
                    </el-col>
                    <el-col :span="12">
                        <el-input v-model="queryObj.description" placeholder="要求描述（或链接）" />
                    </el-col>
                </el-row>
            </el-col>
            <el-col :span="1">
                <el-row>
                    <el-button type="primary" @click="query">查询</el-button>
                </el-row>
            </el-col>
        </el-row>
        <divTable :columnObj="columnObj" :tableData="tableData" :pageObj="pageObj" />
    </div>
</template>

<script setup>
import { reactive } from 'vue'

import divTable from '../../components/divTable/index.vue'

const queryObj = reactive({
    typeUID: "",
    name: "",
    specification: "",
})

const columnObj = {
    headers: [
        {
            prop: "no",
            label: "合同编号",
        },
        {
            prop: "employee.name",
            label: "申请人",
        },
        {
            prop: "startDate",
            label: "申请时间",
        },
        {
            prop: "type",
            label: "类别",
        },
        {
            prop: "product",
            label: "物品",
        },
        {
            prop: "specification",
            label: "型号规格",
        },
        {
            prop: "buyNumber",
            label: "采购数量",
        },
        {
            prop: "unit",
            label: "单位",
        },
        {
            prop: "description",
            label: "要求描述/链接",
        },
        {
            prop: "remarks",
            label: "技术备注",
        },
        {
            prop: "buyDate",
            label: "采购/下单日期",
        },
        {
            prop: "arriveDate",
            label: "到货日期",
        },
        {
            prop: "buyRemarks",
            label: "采购备注",
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