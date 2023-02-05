<template>
    <el-row :gutter="20">
        <el-col :span="6" :offset="8">
            <el-input v-model="base.model.product.name" placeholder="产品名称" clearable maxlength="25" />
        </el-col>
        <el-col :span="1">
            <el-row>
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-row>
        </el-col>
    </el-row>

    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { reactive, onBeforeMount } from 'vue'
import { queryHistoryProducts } from "@/api/history"

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    model: {
        product: {
            name: "",
        }
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "时间",
            },
            {
                prop: "product.name",
                label: "产品",
            },
            {
                prop: "number",
                label: "数量",
            },
            {
                prop: "employee.name",
                label: "员工",
            },
            {
                prop: "remark",
                label: "备注",
            },
        ]
    },
    tableData: [],
    pageData: {
        total: 0,
        pageSize: 10,
        pageNo: 1
    },
    query: () => {
        queryHistoryProducts(base.model, base.pageData).then((res) => {
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
})

onBeforeMount(() => {
    base.query()
})

</script>