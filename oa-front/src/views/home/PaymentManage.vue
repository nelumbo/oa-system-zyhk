<template>
    <div>
        <el-row :gutter="20" align="middle">
            <el-col :span="22">
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-input v-model="queryObj.no" placeholder="合同编号" />
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.employeeName" placeholder="业务员" />
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.companyName" placeholder="客户单位" />
                    </el-col>
                    <el-col :span="6">
                        <el-input v-model="queryObj.customerName" placeholder="客户名称" />
                    </el-col>
                </el-row>
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-select v-model="queryObj.collectionStatus" placeholder="回款状态" clearable
                            style="width: 100%;">
                            <el-option v-for="item in collectionStatusItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
                    </el-col>
                    <el-col :span="6">
                        <el-select v-model="queryObj.invoiceType" placeholder="开票类型" clearable style="width: 100%;">
                            <el-option v-for="item in invoiceTypeItems" :key="item.value" :label="item.label"
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
                <el-row :gutter="20">
                    <el-col :span="6">
                        <el-select v-model="queryObj.isPreDeposit" placeholder="预存款合同" clearable style="width: 100%;">
                            <el-option v-for="item in isPreDepositItems" :key="item.value" :label="item.label"
                                :value="item.value" />
                        </el-select>
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
import { collectionStatusItems, invoiceTypeItems, isPreDepositItems } from '../../utils/magic'

import divTable from '../../components/divTable/index.vue'

const queryObj = reactive({
    no: "",
    employeeName: "",
    companyName: "",
    customerName: "",
    collectionStatus: 1,
    invoiceType: null,
    startDate: new Date().getFullYear() + "-01-01",
    endDate: new Date().getFullYear() + "-12-31",
    isPreDeposit: null,
})

const columnObj = {
    headers: [
        {
            prop: "no",
            label: "合同编号",
        },
        {
            prop: "office.name",
            label: "办事处",
        },
        {
            prop: "employee.name",
            label: "业务员",
        },
        {
            prop: "customer.company.name",
            label: "客户单位",
        },
        {
            prop: "customer.name",
            label: "客户",
        },
        {
            prop: "payType",
            label: "付款类型",
        },
        {
            prop: "isPreDeposit",
            label: "是否预存款",
        },
        {
            prop: "totalAmount",
            label: "总金额",
        },
        {
            prop: "paymentTotalAmount",
            label: "总回款额(CNY)",
        },
        {
            prop: "notPaymentTotalAmount",
            label: "未回款额(CNY)",
        },
        {
            prop: "invoiceType",
            label: "开票类型",
        },
        {
            prop: "endPaymentDate",
            label: "完成回款时间",
        },
        {
            prop: "collectionStatus",
            label: "回款状态",
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