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

    <el-dialog v-model="add.dialogVisible" title="采购" width="75%" :show-close="false">
        <el-divider content-position="left">
            <h2>产品库</h2>
        </el-divider>
        <el-form :model="add.model" label-width="100px">
            <el-row :gutter="20">
                <el-col :span="6" :offset="2">
                    <el-input v-model="add.model.name" placeholder="名称" clearable maxlength="25" />
                </el-col>
                <el-col :span="6">
                    <el-input v-model="add.model.version" placeholder="型号" clearable maxlength="25" />
                </el-col>
                <el-col :span="6">
                    <el-input v-model="add.model.specification" placeholder="规格" clearable maxlength="25" />
                </el-col>
                <el-col :span="1">
                    <el-button type="primary" @click="add.query">查询</el-button>
                </el-col>
            </el-row>
        </el-form>
        <divTable :columnObj="add.column" :tableData="add.tableData" :pageData="add.pageData"
            :handleSizeChange="add.handleSizeChange" :handleCurrentChange="add.handleCurrentChange" />
    </el-dialog>

    <el-dialog v-model="addDLC.dialogVisible" title="添加" width="50%" :show-close="false">
        <el-form :model="addDLC.model" label-width="100px" :rules="rules" ref="addForm">
            <el-form-item label="类型" prop="type">
                <el-select v-model="addDLC.model.type" placeholder="请选择你的采购类型">
                    <el-option v-for="purchasing in PurchasingTypeSelectItems" :label="purchasing.label"
                        :value="purchasing.value" />
                </el-select>
            </el-form-item>
            <el-form-item label="数量" prop="realNumber">
                <el-input-number v-model="addDLC.model.realNumber" :controls="false" :min="1" :max="99999" />
            </el-form-item>
            <el-form-item label="单价" prop="price">
                <el-input-number v-model="addDLC.model.price" :controls="false" :min="0" :max="99999" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="addDLC.submit" :disabled="addDLC.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="check.dialogVisible" title="确认" width="50%" :show-close="false">
        <el-form :model="check.model" label-width="100px" :rules="rules" ref="checkForm">
            <el-form-item label="数量" prop="realNumber">
                <el-input-number v-model="check.model.realNumber" :controls="false" :min="1" :max="99999" />
            </el-form-item>
            <el-form-item label="单价" prop="price">
                <el-input-number v-model="check.model.price" :controls="false" :min="0" :max="99999" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="check.submit" :disabled="addDLC.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="approve.dialogVisible" title="审核" width="50%" :show-close="false">
        <el-form :model="approve.model" label-width="100px">
            <el-form-item label="产品名称">
                <el-input v-model.trim="approve.model.product.name" disabled />
            </el-form-item>
            <el-form-item label="型号">
                <el-input v-model.trim="approve.model.product.verison" disabled />
            </el-form-item>
            <el-form-item label="品牌">
                <el-input v-model.trim="approve.model.product.brand" disabled />
            </el-form-item>
            <el-form-item label="实际采购数量">
                <el-input v-model.trim="approve.model.realNumber" disabled />
            </el-form-item>
            <el-form-item label="参考价格">
                <el-input v-model.trim="approve.model.product.purchasePrice" disabled />
            </el-form-item>
            <el-form-item label="采购价格">
                <el-input v-model.trim="approve.model.price" disabled />
            </el-form-item>
            <el-form-item label="采购总价">
                <el-input v-model.trim="approve.model.totalPrice" disabled />
            </el-form-item>
            <el-divider />
            <el-form-item label="预期到货日期">
                <el-date-picker v-model="approve.model.endDate" type="date" placeholder="结束时间" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="approve.pass" :disabled="approve.submitDisabled"
                        style="margin-right: 250px;">通过</el-button>
                    <el-button type="danger" @click="approve.reject" :disabled="approve.submitDisabled">拒绝</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

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

    <el-dialog v-model="final.dialogVisible" title="完成确认" width="50%" :show-close="false">
        <h1>是否确定该条采购记录已经记录？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="final.submit" :disabled="final.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { message } from '@/components/divMessage/index'
import { PurchasingTypeSelectItems } from '@/utils/magic'
import { addPurchasing, approvePurchasing, finalPurchasingProduct, finalPurchasing, queryPurchasings } from "@/api/purchasing"
import { queryProducts } from "@/api/product"
import { reg_number, reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const checkForm = ref(null)

const rules = reactive({
    type: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    realNumber: [
        { required: true, pattern: reg_number, message: '请输入大于零的有效数字', trigger: 'blur' },
    ],
    price: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
})

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
                            if (row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "确认",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openCheckDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 2) {
                                return true
                            }
                            return false
                        },
                        label: "审核",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openApproveDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 3 && row.productStatus == 1) {
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
                    {
                        isShow: (index, row) => {
                            if (row.status == 3 && row.productStatus == 2 && row.payStatus == 2 && row.invoiceStatus == 2) {
                                return true
                            }
                            return false
                        },
                        label: "完成",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openFinalDialog(index, row)
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
    openAddDialog() {
        add.query()
        add.dialogVisible = true
    },
    openCheckDialog(index, row) {
        check.model.id = row.id
        check.dialogVisible = true
    },
    openApproveDialog: (index, row) => {
        approve.model = row
        approve.dialogVisible = true
    },
    openFinalProductDialog: (index, row) => {
        finalProduct.model.id = row.id
        finalProduct.dialogVisible = true
    },
    openFinalDialog: (index, row) => {
        final.model.id = row.id
        final.dialogVisible = true
    },
})

const add = reactive({
    dialogVisible: false,
    model: {
        name: "",
        version: "",
        specification: "",
    },
    column: {
        headers: [
            {
                prop: "type.name",
                label: "类型",
                width: "5%",
            },
            {
                prop: "name",
                label: "名称",
                width: "10%",
            },
            {
                prop: "version",
                label: "型号",
                width: "8%",
            },
            {
                prop: "brand",
                label: "品牌",
                width: "5%",
            },
            {
                prop: "specification",
                label: "规格",
                width: "10%",
            },
            {
                prop: "number",
                label: "可售数量",
                width: "10%",
            },
            {
                prop: "unit",
                label: "单位",
                width: "5%",
            },
            {
                prop: "attribute.standardPrice",
                label: "标准售价(元)",
                width: "10%",
            },
            {
                prop: "attribute.standardPriceUSD",
                label: "标准售价(美元)",
                width: "10%",
            },
            {
                type: "boolean",
                prop: "isFree",
                label: "小零配件",
                width: "7%",
            },
            {
                type: "operation",
                label: "操作",
                width: "10%",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "采购",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => add.openAddDLCDialog(index, row)
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
        queryProducts(add.model, add.pageData).then((res) => {
            if (res.status == 1) {
                add.tableData = res.data.data
                add.pageData.total = res.data.total
                add.pageData.pageSize = res.data.pageSize
                add.pageData.pageNo = res.data.pageNo
            } else {
                message("查询失败", "error")
            }
        })
    },
    openAddDLCDialog(index, row) {
        addDLC.model.productID = row.id
        addDLC.dialogVisible = true
    },
})

const addDLC = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        productID: null,
        type: null,
        price: 0,
        realNumber: 0,
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                addDLC.submitDisabled = true
                addPurchasing(addDLC.model).then((res) => {
                    if (res.status == 1) {
                        message("录入成功", "success")
                        base.query()
                    } else {
                        message("录入失败", "error")
                    }
                    addDLC.dialogVisible = false
                    addDLC.model = {
                        productID: null,
                        type: null,
                        price: 0,
                        realNumber: 0,
                    }
                    addDLC.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

const check = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        realNumber: 0,
        price: 0,
    },
    submit: () => {
        checkForm.value.validate((valid) => {
            if (valid) {
                check.submitDisabled = true
                approvePurchasing(check.model).then((res) => {
                    if (res.status == 1) {
                        message("审核成功", "success")
                        base.query()
                    } else {
                        message("审核失败", "error")
                    }
                    check.dialogVisible = false
                    check.model = {
                        id: null,
                        realNumber: 0,
                        price: 0,
                    }
                    check.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const approve = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        isPass: null,
        realNumber: 0,
        price: 0,
        totalPrice: 0,
        endDate: "",
        product: {
            purchasePrice: 0,
            name: "",
            verison: "",
            brand: "",
        }
    },
    submit: () => {
        approve.submitDisabled = true
        approvePurchasing(
            {
                "id": approve.model.id,
                "isPass": approve.model.isPass,
                "endDate": approve.model.endDate
            }
        ).then((res) => {
            if (res.status == 1) {
                message("审核成功", "success")
                base.query()
            } else {
                message("审核失败", "error")
            }
            approve.dialogVisible = false
            approve.model = {
                id: null,
                isPass: null,
                product: {
                    name: "",
                }
            }
            approve.submitDisabled = false
        })
    },
    pass: () => {
        approve.model.isPass = true
        approve.submit()
    },
    reject: () => {
        approve.model.isPass = false
        approve.submit()
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

const final = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
    },
    submit: () => {
        final.submitDisabled = true
        finalPurchasing(final.model).then((res) => {
            if (res.status == 1) {
                message("操作成功", "success")
                base.query()
            } else {
                message("操作失败", "error")
            }
            final.dialogVisible = false
            final.model = {
                id: null,
            }
            final.submitDisabled = false
        })
    }
})

onBeforeMount(() => {
    base.query()
})
</script>