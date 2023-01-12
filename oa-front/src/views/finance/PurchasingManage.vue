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

    <el-dialog v-model="finalPay.dialogVisible" title="确认" width="50%" :show-close="false">
        <h1>是否确定该条采购记录已经付款？</h1>
        <el-form :model="finalPay.model" label-width="60px">
            <el-form-item label="备注">
                <el-input v-model="finalPay.model.payRemark" type="textarea" :autosize="{ minRows: 9, maxRows: 18 }"
                    maxlength="300" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="finalPay.submit"
                        :disabled="finalPay.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
    <el-dialog v-model="finalInvoice.dialogVisible" title="确认" width="50%" :show-close="false">
        <h1>是否确定该条采购记录已经收到发票？</h1>
        <el-form :model="finalInvoice.model" label-width="60px">
            <el-form-item label="备注">
                <el-input v-model="finalInvoice.model.invoiceRemark" type="textarea"
                    :autosize="{ minRows: 9, maxRows: 18 }" maxlength="300" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="finalInvoice.submit"
                        :disabled="finalInvoice.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { message } from '@/components/divMessage/index'
import { finalPurchasingPay, finalPurchasingInvoice, queryPurchasings } from "@/api/purchasing"

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
                prop: "no",
                label: "编号",
                width: "15%"
            },
            {
                prop: "createDate",
                label: "发起时间",
                width: "10%"

            },
            {
                prop: "employee.name",
                label: "发起人",
                width: "10%"
            },
            {
                prop: "product.name",
                label: "名称",
                width: "10%"
            },
            {
                prop: "number",
                label: "数量",
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
                width: "20%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == 3 && row.payStatus == 1) {
                                return true
                            }
                            return false
                        },
                        label: "付款确认",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openFinalPayDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 3 && row.invoiceStatus == 1) {
                                return true
                            }
                            return false
                        },
                        label: "发票确认",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openFinalInvoiceDialog(index, row)
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
    openFinalPayDialog: (index, row) => {
        finalPay.model.id = row.id
        finalPay.dialogVisible = true
    },
    openFinalInvoiceDialog: (index, row) => {
        finalInvoice.model.id = row.id
        finalInvoice.dialogVisible = true
    },
})

const finalPay = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        payRemark: "",
    },
    submit: () => {
        finalPay.submitDisabled = true
        finalPurchasingPay(finalPay.model).then((res) => {
            if (res.status == 1) {
                message("操作成功", "success")
                base.query()
            } else {
                message("操作失败", "error")
            }
            finalPay.dialogVisible = false
            finalPay.model = {
                id: null,
                remark: "",
            }
            finalPay.submitDisabled = false
        })
    }
})

const finalInvoice = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        invoiceRemark: "",
    },
    submit: () => {
        finalInvoice.submitDisabled = true
        finalPurchasingInvoice(finalInvoice.model).then((res) => {
            if (res.status == 1) {
                message("操作成功", "success")
                base.query()
            } else {
                message("操作失败", "error")
            }
            finalInvoice.dialogVisible = false
            finalInvoice.model = {
                id: null,
                remark: "",
            }
            finalInvoice.submitDisabled = false
        })
    }
})

onBeforeMount(() => {
    base.query()
})
</script>