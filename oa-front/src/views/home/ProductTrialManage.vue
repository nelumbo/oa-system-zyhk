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
        <el-col :span="1">
            <el-button type="success" @click="base.openAddDialog">发起</el-button>
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

    <el-dialog v-model="add.dialogVisible" title="发起" width="75%" :show-close="false">
        <el-row :gutter="20">
            <el-col :span="6" :offset="2">
                <el-select v-model="add.model.typeID" placeholder="类型" clearable style="width: 100%;">
                    <el-option v-for="item in add.productTypes" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="add.model.name" placeholder="产品名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="6">
                <el-input v-model="add.model.specification" placeholder="规格" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="add.query">查询</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="add.column" :tableData="add.tableData" :pageData="add.pageData"
            :handleSizeChange="add.handleSizeChange" :handleCurrentChange="add.handleCurrentChange" />
    </el-dialog>

    <el-dialog v-model="addView.dialogVisible" title="产品查看" width="50%" :show-close="false">
        <el-form :model="addView.model" label-width="150px">
            <el-form-item label="类型">
                <el-input v-model.trim="addView.model.type.name" disabled />
            </el-form-item>
            <el-form-item label="名称">
                <el-input v-model.trim="addView.model.name" disabled />
            </el-form-item>
            <el-form-item label="品牌">
                <el-input v-model.trim="addView.model.brand" disabled />
            </el-form-item>
            <el-form-item label="规格">
                <el-input v-model.trim="addView.model.specification" disabled />
            </el-form-item>
            <el-form-item label="供应商">
                <el-row>
                    <el-col :span="24" v-for="supplier in addView.model.suppliers" :key="supplier.id">
                        {{ supplier.name }}
                    </el-col>
                </el-row>
            </el-form-item>
            <el-form-item label="采购价格(元)">
                <el-input v-model.trim="addView.model.attribute.purchasePrice" disabled />
            </el-form-item>
            <el-form-item label="采购价格(美元)">
                <el-input v-model.trim="addView.model.attribute.purchasePriceUSD" disabled />
            </el-form-item>
            <el-form-item label="标准售价(元)">
                <el-input v-model.trim="addView.model.attribute.standardPrice" disabled />
            </el-form-item>
            <el-form-item label="标准售价(美元)">
                <el-input v-model.trim="addView.model.attribute.standardPriceUSD" disabled />
            </el-form-item>
            <el-form-item label="库存数量" prop="number">
                <el-input v-model.trim="addView.model.number" disabled />
            </el-form-item>
            <el-form-item label="库存单位" prop="unit">
                <el-input v-model.trim="addView.model.unit" disabled />
            </el-form-item>
            <el-form-item label="供货周期">
                <el-input v-model.trim="addView.model.deliveryCycle" disabled />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="addView.model.remark" type="textarea" :autosize=true disabled />
            </el-form-item>
            <el-form-item label="小零配件">
                <div v-if="addView.model.isFree">
                    是
                </div>
                <div v-if="!addView.model.isFree">
                    否
                </div>
            </el-form-item>
        </el-form>
    </el-dialog>

    <el-dialog v-model="addDLC.dialogVisible" title="数量&备注" width="50%" :show-close="false">
        <el-form :model="addDLC.model" label-width="60px" :rules="rules" ref="addForm">
            <el-form-item label="数量" prop="number">
                <el-input-number v-model="addDLC.model.number" :controls="false" :min="1" />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="addDLC.model.remark1" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
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

    <el-dialog v-model="del.dialogVisible" title="删除" width="50%" :show-close="false">
        <h1>是否确定删除该条记录？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="del.submit" :disabled="del.submitDisabled">确定</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { ref, reactive, onBeforeMount, computed } from 'vue'
import { ProductTrialStatusItems } from '@/utils/magic'
import { queryMyProductTrials } from "@/api/my"
import { queryProduct, queryProducts } from "@/api/product"
import { queryAllProductType } from "@/api/product_type"
import { addProductTrial, delProductTrial } from "@/api/product_trial"
import { message } from '@/components/divMessage/index'
import { reg_number } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const rules = reactive({
    productID: [
        { required: true, message: '请选择产品', trigger: 'blur' },
    ],
    number: [
        { required: true, pattern: reg_number, message: '请输入有效数字', trigger: 'blur' }
    ],
})

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
                            if (row.status == -1 || row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "删除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openDelDialog(index, row)
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
        queryMyProductTrials(base.model, base.pageData).then((res) => {
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
        add.query()
        add.dialogVisible = true
    },
    openViewDialog: (index, row) => {
        view.model = row
        view.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.dialogVisible = true
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

const add = reactive({
    dialogVisible: false,
    productTypes: [],
    model: {
        typeID: "",
        name: "",
        specification: "",
    },
    column: {
        headers: [
            {
                prop: "type.name",
                label: "类型",
                width: "10%",
            },
            {
                prop: "name",
                label: "名称",
                width: "20%",
            },
            {
                prop: "brand",
                label: "品牌",
                width: "15%",
            },
            {
                prop: "specification",
                label: "规格",
                width: "15%",
            },
            {
                prop: "number",
                label: "库存数量",
                width: "10%",
            },
            {
                prop: "unit",
                label: "单位",
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
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => add.openAddViewDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "添加",
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
    handleSizeChange: (e) => {
        add.pageData.pageSize = e
        add.pageData.pageNo = 1
        add.query()
    },
    handleCurrentChange: (e) => {
        add.pageData.pageNo = e
        add.query()
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
    openAddViewDialog: (index, row) => {
        queryProduct(row.id).then((res) => {
            if (res.status == 1) {
                addView.model = res.data
            }
        })
        addView.dialogVisible = true
    },
    openAddDLCDialog: (index, row) => {
        addDLC.model.productID = row.id
        addDLC.model.number = 1
        addDLC.dialogVisible = true
    },
})

const addView = reactive({
    dialogVisible: false,
    model: {
        type: {
            name: ""
        },
        name: "",
        brand: "",
        specification: "",
        suppliers: [],
        attribute: {
            purchasePrice: 0,
            purchasePriceUSD: 0,
            standardPrice: 0,
            standardPriceUSD: 0,
        },
        number: 0,
        unit: "",
        deliveryCycle: "",
        remark: "",
        isFree: false,
    }
})

const addDLC = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        productID: "",
        number: 1,
        remark1: "",
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                addDLC.submitDisabled = true
                addProductTrial(addDLC.model).then((res) => {
                    if (res.status == 1) {
                        message("录入成功", "success")
                        base.query()
                    } else {
                        message("录入失败", "error")
                    }
                    add.dialogVisible = false
                    addDLC.dialogVisible = false
                    addDLC.model = {
                        productID: "",
                        number: 1,
                        remark1: "",
                    }
                    addDLC.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

const del = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
    },
    submit: () => {
        del.submitDisabled = true
        delProductTrial(del.model.id).then((res) => {
            if (res.status == 1) {
                message("删除成功", "success")
                base.query()
            } else {
                message("删除失败", "error")
            }
            del.dialogVisible = false
            del.model = {
                id: 0,
            }
            del.submitDisabled = false
        })
    }
})

onBeforeMount(() => {
    base.query()
    queryAllProductType().then((res) => {
        if (res.status == 1) {
            add.productTypes = res.data
        }
    })
})
</script>