<template>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { message } from '@/components/divMessage/index'
import { queryProductCalls } from "@/api/product_call"

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    column: {
        headers: [
            {
                prop: "createDate",
                label: "产品",
            },
            {
                prop: "product.name",
                label: "产品",
            },
            {
                prop: "product.number",
                label: "库存数量",
            },
            {
                prop: "product.callNumber",
                label: "报警数量",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "采购",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        // onClick: (index, row) => base.openCheckDialog(index, row)
                    },
                ]
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
        queryProductCalls(null, base.pageData).then((res) => {
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
        base.pageData.pageSize = ae
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