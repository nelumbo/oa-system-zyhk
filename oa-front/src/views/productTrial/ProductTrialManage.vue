<template>
    <el-row :gutter="20">
        <el-col :span="6" :offset="8">
            <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                <el-option v-for="item in ProductTrialStatusItems" :key="item.value" :label="item.label"
                    :value="item.value" />
            </el-select>
        </el-col>
        <el-col :span="1">
            <el-button type="primary" @click="base.query">查询</el-button>
        </el-col>
    </el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

    <el-dialog v-model="view.dialogVisible" title="查看" width="50%" :show-close="false">
        <el-form :model="view.model" label-width="100px">
            <el-form-item label="创建日期">
                <el-input v-model.trim="view.model.createDate" readonly />
            </el-form-item>
            <el-form-item label="业务员">
                <el-input v-model.trim="view.model.employee.name" readonly />
            </el-form-item>
            <el-form-item label="产品">
                <el-input v-model.trim="view.model.product.name" readonly />
            </el-form-item>
            <el-form-item label="审核日期">
                <el-input v-model.trim="view.model.auditDate" readonly />
            </el-form-item>
            <el-form-item label="发货日期">
                <el-input v-model.trim="view.model.shipmentDate" readonly />
            </el-form-item>
            <el-form-item label="归还日期">
                <el-input v-model.trim="view.model.inventoryDate" readonly />
            </el-form-item>
            <el-form-item label="完成日期">
                <el-input v-model.trim="view.model.finalDate" readonly />
            </el-form-item>
            <el-form-item label="业务员备注">
                <el-input v-model.trim="view.model.remark1" type="textarea" autosize readonly />
            </el-form-item>
            <el-form-item label="发货人员备注">
                <el-input v-model.trim="view.model.remark3" type="textarea" autosize readonly />
            </el-form-item>
            <el-form-item label="收货人员备注">
                <el-input v-model.trim="view.model.remark4" type="textarea" autosize readonly />
            </el-form-item>
            <el-form-item label="完成人员备注">
                <el-input v-model.trim="view.model.remark5" type="textarea" autosize readonly />
            </el-form-item>
            <el-form-item label="状态">
                <el-input v-model.trim="view.statusString" readonly />
            </el-form-item>
        </el-form>
    </el-dialog>

    <el-dialog v-model="approve.dialogVisible" title="审核" width="50%" :show-close="false">
        <el-form :model="approve.model" label-width="100px">
            <el-form-item label="产品">
                <el-input v-model.trim="approve.productTrial.product.name" readonly />
            </el-form-item>
            <el-form-item label="数量">
                <el-input v-model.trim="approve.productTrial.number" readonly />
            </el-form-item>
            <el-form-item label="业务员备注">
                <el-input v-model.trim="approve.productTrial.remark1" type="textarea" autosize readonly />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="approve.pass" :disabled="approve.submitDisabled"
                        style="margin-right: 250px;">通过</el-button>
                    <el-button type="danger" @click="approve.reject" :disabled="approve.submitDisabled">驳回</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="next.dialogVisible" title="提交" width="50%" :show-close="false">
        <el-form :model="next.model" label-width="60px">
            <el-form-item label="备注" v-if="next.model.status == 2">
                <el-input v-model.trim="next.model.remark3" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
            </el-form-item>
            <el-form-item label="备注" v-if="next.model.status == 3">
                <el-input v-model.trim="next.model.remark4" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
            </el-form-item>
            <el-form-item label="备注" v-if="next.model.status == 4">
                <el-input v-model.trim="next.model.remark5" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="next.submit" :disabled="next.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { reactive, onBeforeMount, computed } from 'vue'
import { ProductTrialStatusItems } from '@/utils/magic'
import { nextProductTrial, queryProductTrials } from "@/api/product_trial"
import { message } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    model: {
        status: null,
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "录入日期",
                width: "10%",
            },
            {
                prop: "employee.name",
                label: "业务员",
                width: "10%",
            },
            {
                prop: "product.name",
                label: "产品",
                width: "15%",
            },
            {
                prop: "number",
                label: "数量",
                width: "5%",
            },
            {
                prop: "auditor.name",
                label: "审核员",
                width: "10%",
            },
            {
                prop: "shipment.name",
                label: "物流人员",
                width: "10%",
            },
            {
                prop: "inventory.name",
                label: "收件人员",
                width: "10%",
            },
            {
                prop: "final.name",
                label: "确认人员",
                width: "10%",
            },
            {
                type: "transform",
                prop: "status",
                items: ProductTrialStatusItems,
                label: "状态",
                width: "5%",
            },
            {
                type: "operation",
                label: "操作",
                width: "15%",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "查看",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openViewDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "审批",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openApproveDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status > 1 && row.status < 5) {
                                return true
                            }
                            return false
                        },
                        label: "下一步",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openNextDelDialog(index, row)
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
        queryProductTrials(base.model, base.pageData).then((res) => {
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
    openViewDialog: (index, row) => {
        view.model = row
        view.dialogVisible = true
    },
    openApproveDialog: (index, row) => {
        approve.productTrial = row
        approve.model.id = row.id
        approve.dialogVisible = true
    },
    openNextDelDialog: (index, row) => {
        next.model.id = row.id
        next.model.status = row.status
        next.dialogVisible = true
    }
})

const view = reactive({
    dialogVisible: false,
    statusString: computed(() => {
        var temp = "";
        ProductTrialStatusItems.some((item) => {
            if (item.value == view.model.status) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    model: {
        product: {
            name: "",
        },
        number: 0,
        employee: {
            name: "",
        },
        auditor: {
            name: "",
        },
        shipment: {
            name: "",
        },
        inventory: {
            name: "",
        },
        final: {
            name: "",
        },
        createDate: "",
        auditDate: "",
        shipmentDate: "",
        inventoryDate: "",
        finalDate: "",
        remark1: "",
        remark3: "",
        remark4: "",
        remark5: "",
        status: null,
    }
})

const approve = reactive({
    dialogVisible: false,
    submitDisabled: false,
    productTrial: {
        product: {
            name: "",
        },
        number: 0,
        reamrk1: "",
    },
    model: {
        id: null,
        status: null,
    },
    submit: () => {
        approve.submitDisabled = true
        nextProductTrial(approve.model).then((res) => {
            if (res.status == 1) {
                message("审核成功", "success")
                base.query()
            } else {
                message("审核失败", "error")
            }
            approve.dialogVisible = false
            approve.productTrial = {
                product: {
                    name: "",
                },
                number: 0,
                reamrk1: "",
            }
            approve.model = {
                id: null,
                status: null,
            }
            approve.submitDisabled = false
        })
    },
    pass: () => {
        approve.model.status = 2
        approve.submit()
    },
    reject: () => {
        approve.model.status = -1
        approve.submit()
    },
})

const next = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        status: null,
        remark3: "",
        remark4: "",
        remark5: "",
    },
    submit: () => {
        next.submitDisabled = true
        next.model.status = next.model.status + 1
        nextProductTrial(next.model).then((res) => {
            if (res.status == 1) {
                message("提交成功", "success")
                base.query()
            } else {
                message("提交失败", "error")
            }
            next.dialogVisible = false
            next.model = {
                id: null,
                status: null,
                remark3: "",
                remark4: "",
                remark5: "",
            }
            next.submitDisabled = false
        })
    },
})

onBeforeMount(() => {
    base.query()
})
</script>