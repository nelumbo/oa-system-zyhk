<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-input v-model="base.model.no" placeholder="合同编号" clearable maxlength="50" />
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.regionID" placeholder="省份" clearable style="width: 100%;">
                    <el-option v-for="item in base.model.regions" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.customer.customerCompany.name" placeholder="客户单位" clearable
                    maxlength="25" />
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.customer.name" placeholder="客户名称" clearable maxlength="25" />
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-select v-model="base.model.status" placeholder="合同状态" clearable style="width: 100%;">
                    <el-option v-for="item in contractStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.productionStatus" placeholder="生产状态" clearable style="width: 100%;">
                    <el-option v-for="item in productionStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.collectionStatus" placeholder="回款状态" clearable style="width: 100%;">
                    <el-option v-for="item in collectionStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.payType" placeholder="付款类型" clearable style="width: 100%;">
                    <el-option v-for="item in payTypeItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
            <el-col :span="1">
                <el-button type="success" @click="base.openAddDialog">录入</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-date-picker v-model="base.model.startDate" type="date" placeholder="开始时间" style="width: 100%;" />
            </el-col>
            <el-col :span="5">
                <el-date-picker v-model="base.model.endDate" type="date" placeholder="结束时间" style="width: 100%;" />
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.isSpecialNum" placeholder="特殊合同" clearable style="width: 100%;">
                    <el-option v-for="item in boolItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.isPreDepositNum" placeholder="预存款合同" clearable style="width: 100%;">
                    <el-option v-for="item in boolItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />


    </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { ref, reactive, onBeforeMount } from 'vue'
import { contractStatusItems, productionStatusItems, collectionStatusItems, payTypeItems, boolItems, isSpecialItems, isPreDepositItems } from '@/utils/magic'
import { queryMyContracts } from "@/api/my"
import { delContract, queryContract } from "@/api/contract"
import { message } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const router = useRouter()
const editForm = ref(null)
const rules = reactive({
})

const base = reactive({
    regions: [],
    model: {
        no: "",
        regionID: null,
        customer: {
            name: "",
            customerCompany: {
                name: "",
            }
        },
        status: null,
        productionStatus: null,
        collectionStatus: null,
        payType: null,
        startDate: new Date().getFullYear() + "-01-01",
        endDate: new Date().getFullYear() + "-12-31",
        isSpecialNum: null,
        isPreDepositNum: null,
    },
    column: {
        headers: [
            {
                prop: "no",
                label: "合同编号",
            },
            {
                prop: "customer.customerCompany.name",
                label: "客户单位",
            },
            {
                prop: "customer.name",
                label: "客户",
            },
            {
                prop: "estimatedDeliveryDate",
                label: "合同交货日期",
            },
            {
                prop: "endDeliveryDate",
                label: "实际交货日期",
            },
            {
                prop: "preDeposit",
                label: "剩余预存款金额",
            },
            {
                prop: "totalAmount",
                label: "总金额",
            },
            {
                prop: "notPaymentTotalAmount",
                label: "未回款额(CNY)",
            },
            {
                prop: "status",
                label: "状态",
            },
            {
                type: "operation",
                label: "操作",
                operations: [{
                    type: "primary",
                    label: "查看",
                    icon: "",
                    color: '',
                    buttonClick: () => {
                        dialogType.viewType = true
                    },
                    isShow: () => {
                        return true;
                    }
                }]
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
        queryMyContracts(base.model, base.pageData).then((res) => {
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
        router.push("entry")
    },
    openViewDialog: (index, row) => {
        view.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.dialogVisible = true
    }
})

onBeforeMount(() => {
    base.query()
})
</script>