<template>
    <el-row :gutter="20">
        <el-col :span="6" :offset="2">
            <el-input v-model="base.model.supplierName" placeholder="供应商名称" clearable maxlength="25" />
        </el-col>
        <el-col :span="6">
            <el-date-picker v-model="base.model.startDate" type="date" placeholder="开始时间" style="width: 100%;" />
        </el-col>
        <el-col :span="6">
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

    <el-dialog v-model="finalProduct.dialogVisible" title="确认" width="50%" :show-close="false">
        <h1>是否确定该条采购记录已经收货？</h1>
        <el-form :model="finalProduct.model" label-width="60px">
            <el-form-item label="备注">
                <el-input v-model="finalProduct.model.productRemark" type="textarea"
                    :autosize="{ minRows: 3, maxRows: 6 }" maxlength="300" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="finalProduct.submit"
                        :disabled="finalProduct.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { reactive, onBeforeMount } from 'vue'
import { message } from '@/components/divMessage/index'
import { finalPurchasingProduct, queryPurchasings } from "@/api/purchasing"

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    model: {
        supplierName: "",
        startDate: "",
        endDate: "",
    },
    column: {
        headers: [
            {
                type: "purchasingNo",
                prop: "no",
                label: "编号",
                width: "10%"
            },
            {
                prop: "createDate",
                label: "发起时间",
                width: "5%"
            },
            {
                prop: "endDate",
                label: "预期时间",
                width: "5%"
            },
            {
                prop: "employee.name",
                label: "发起人",
                width: "5%"
            },
            {
                prop: "product.name",
                label: "名称",
                width: "8%"
            },
            {
                prop: "product.version",
                label: "型号",
                width: "7%"
            },
            {
                prop: "product.specification",
                label: "规格",
                width: "10%"
            },
            {
                prop: "number",
                label: "需求数量",
                width: "5%"
            },
            {
                prop: "realNumber",
                label: "实际数量",
                width: "5%"
            },
            {
                prop: "product.unit",
                label: "单位",
                width: "5%"
            },
            {
                prop: "price",
                label: "单价",
                width: "5%"
            },
            {
                prop: "totalPrice",
                label: "总价",
                width: "5%"
            },
            {
                type: "purchasingStatus",
                prop: "status",
                label: "状态",
                width: "15%"
            },
            {
                type: "operation",
                label: "操作",
                width: "10%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == 3 && row.productStatus == 1 && user().my.pids.includes('100')) {
                                return true
                            }
                            return false
                        },
                        label: "确认收货",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openFinalProductDialog(index, row)
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
        queryPurchasings(base.model, base.pageData).then((res) => {
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
    openFinalProductDialog: (index, row) => {
        finalProduct.model.id = row.id
        finalProduct.dialogVisible = true
    },
})

const finalProduct = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        productRemark: "",
    },
    submit: () => {
        finalProduct.submitDisabled = true
        finalPurchasingProduct(finalProduct.model).then((res) => {
            if (res.status == 1) {
                message("操作成功", "success")
                base.query()
            } else {
                message("操作失败", "error")
            }
            finalProduct.dialogVisible = false
            finalProduct.model = {
                id: null,
                productRemark: "",
            }
            finalProduct.submitDisabled = false
        })
    }
})


onBeforeMount(() => {
    base.query()
})
</script>