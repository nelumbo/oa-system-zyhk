<template>
    <div>
        <el-row :gutter="10">
            <el-col :span="6" :offset="2">
                <el-select v-model="base.model.typeID" placeholder="类型" clearable style="width: 100%;">
                    <el-option v-for="item in base.productTypes" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.name" placeholder="产品名称" />
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.specification" placeholder="规格" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
            <el-col :span="1">
                <el-button type="success" @click="base.openAddDialog">添加</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="add.dialogVisible" title="产品添加" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="120px" :rules="rules" ref="addForm">
                <el-form-item label="类型" prop="typeID">
                    <el-select v-model="add.model.typeID" clearable>
                        <el-option v-for="productType in base.productTypes" :key="productType.id"
                            :label="productType.name" :value="productType.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="add.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="品牌">
                    <el-input v-model.trim="add.model.brand" maxlength="50" />
                </el-form-item>
                <el-form-item label="规格">
                    <el-input v-model.trim="add.model.specification" maxlength="50" />
                </el-form-item>
                <el-form-item label="供应商">
                    <el-col :span="24">
                        <el-button type="success" @click="base.openAddDLCDialog">选择供应商</el-button>
                    </el-col>
                    <el-col :span="24" v-for="supplier in add.model.suppliers" :key="supplier.id">
                        {{ supplier.name }}
                    </el-col>
                </el-form-item>
                <el-form-item label="采购价格(元)" prop="attribute.purchasePrice">
                    <el-input-number v-model="add.model.attribute.purchasePrice" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
                <el-form-item label="采购价格(美元)" prop="attribute.purchasePriceUSD">
                    <el-input-number v-model="add.model.attribute.purchasePriceUSD" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
                <el-form-item label="标准价格(元)" prop="attribute.standardPrice">
                    <el-input-number v-model="add.model.attribute.standardPrice" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
                <el-form-item label="标准价格(美元)" prop="attribute.standardPriceUSD">
                    <el-input-number v-model="add.model.attribute.standardPriceUSD" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
                <el-form-item label="库存数量" prop="numberCount">
                    <el-input-number v-model="add.model.numberCount" :controls="false" :min="-10000" :max="99999" />
                </el-form-item>
                <el-form-item label="库存单位" prop="unit">
                    <el-input v-model.trim="add.model.unit" maxlength="50" />
                </el-form-item>
                <el-form-item label="供货周期">
                    <el-input v-model.trim="add.model.deliveryCycle" maxlength="50" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="add.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
                </el-form-item>
                <el-form-item label="小零配件">
                    <el-radio-group v-model="add.model.isFree">
                        <el-radio :label="true">是</el-radio>
                        <el-radio :label="false">否</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="add.submit" :disabled="add.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="addDLC.dialogVisible" title="供应商选择" width="50%" :show-close="false">
            <el-row :gutter="10">
                <el-col :span="6" :offset="2">
                    <el-input v-model="addDLC.model.name" placeholder="供应商名称" clearable maxlength="25" />
                </el-col>
                <el-col :span="6">
                    <el-input v-model="addDLC.model.linkman" placeholder="联系人" clearable maxlength="25" />
                </el-col>
                <el-col :span="6">
                    <el-input v-model="addDLC.model.phone" placeholder="联系电话" clearable maxlength="25" />
                </el-col>
                <el-col :span="1">
                    <el-button type="primary" @click="addDLC.query">查询</el-button>
                </el-col>
            </el-row>
            <divSelectTable :queryObj="addDLC.model" :queryFunc="querySuppliers" :selectObj="add.model.suppliers"
                ref="supplierSelect"></divSelectTable>
            <el-row>
                <el-col :span="24">
                    <h3>已选择：</h3>
                </el-col>
                <el-col :span="22" :offset="2" v-for="supplier in add.model.suppliers" :key="supplier.id">
                    {{ supplier.name }}
                </el-col>
            </el-row>
        </el-dialog>

        <el-dialog v-model="view.dialogVisible" title="产品查看" width="50%" :show-close="false">
            <el-form :model="view.model" label-width="150px">
                <el-form-item label="类型">
                    <el-input v-model.trim="view.model.type.name" disabled />
                </el-form-item>
                <el-form-item label="名称">
                    <el-input v-model.trim="view.model.name" disabled />
                </el-form-item>
                <el-form-item label="品牌">
                    <el-input v-model.trim="view.model.brand" disabled />
                </el-form-item>
                <el-form-item label="规格">
                    <el-input v-model.trim="view.model.specification" disabled />
                </el-form-item>
                <el-form-item label="供应商">
                    <el-row>
                        <el-col :span="24" v-for="supplier in view.model.suppliers" :key="supplier.id">
                            {{ supplier.name }}
                        </el-col>
                    </el-row>
                </el-form-item>
                <el-form-item label="采购价格(元)">
                    <el-input v-model.trim="view.model.attribute.purchasePrice" disabled />
                </el-form-item>
                <el-form-item label="采购价格(美元)">
                    <el-input v-model.trim="view.model.attribute.purchasePriceUSD" disabled />
                </el-form-item>
                <el-form-item label="标准售价(元)">
                    <el-input v-model.trim="view.model.attribute.standardPrice" disabled />
                </el-form-item>
                <el-form-item label="标准售价(美元)">
                    <el-input v-model.trim="view.model.attribute.standardPriceUSD" disabled />
                </el-form-item>
                <el-form-item label="库存数量" prop="numberCount">
                    <el-input v-model.trim="view.model.numberCount" disabled />
                </el-form-item>
                <el-form-item label="库存单位" prop="unit">
                    <el-input v-model.trim="view.model.unit" disabled />
                </el-form-item>
                <el-form-item label="供货周期">
                    <el-input v-model.trim="view.model.deliveryCycle" disabled />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="view.model.remark" type="textarea" :autosize=true disabled />
                </el-form-item>
                <el-form-item label="小零配件">
                    <div v-if="view.model.isFree">
                        是
                    </div>
                    <div v-if="!view.model.isFree">
                        否
                    </div>
                </el-form-item>
            </el-form>
        </el-dialog>

        <el-dialog v-model="edit.baseDialogVisible" title="产品基础信息编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="150px" :rules="rules" ref="editForm">
                <el-form-item label="类型" prop="typeID">
                    <el-select v-model="edit.model.typeID" clearable>
                        <el-option v-for="productType in base.productTypes" :key="productType.id"
                            :label="productType.name" :value="productType.id" />
                    </el-select>
                </el-form-item>
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="edit.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="品牌">
                    <el-input v-model.trim="edit.model.brand" maxlength="50" />
                </el-form-item>
                <el-form-item label="规格">
                    <el-input v-model.trim="edit.model.specification" maxlength="50" />
                </el-form-item>
                <el-form-item label="库存单位" prop="unit">
                    <el-input v-model.trim="edit.model.unit" maxlength="50" />
                </el-form-item>
                <el-form-item label="供应商">
                    <el-col :span="24">
                        <el-button type="success" @click="base.openEditDLCDialog">选择供应商</el-button>
                    </el-col>
                    <el-col :span="24" v-for="supplier in edit.model.suppliers" :key="supplier.id">
                        {{ supplier.name }}
                    </el-col>
                </el-form-item>
                <el-form-item label="供货周期">
                    <el-input v-model.trim="edit.model.deliveryCycle" maxlength="50" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="edit.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
                </el-form-item>
                <el-form-item label="小零配件">
                    <el-radio-group v-model="edit.model.isFree">
                        <el-radio :label="true">是</el-radio>
                        <el-radio :label="false">否</el-radio>
                    </el-radio-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="edit.submitBase"
                            :disabled="edit.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="editDLC.dialogVisible" title="供应商选择" width="50%" :show-close="false">
            <el-row :gutter="10">
                <el-col :span="6" :offset="2">
                    <el-input v-model="editDLC.model.name" placeholder="供应商名称" clearable maxlength="25" />
                </el-col>
                <el-col :span="6">
                    <el-input v-model="editDLC.model.linkman" placeholder="联系人" clearable maxlength="25" />
                </el-col>
                <el-col :span="6">
                    <el-input v-model="editDLC.model.phone" placeholder="联系电话" clearable maxlength="25" />
                </el-col>
                <el-col :span="1">
                    <el-button type="primary" @click="editDLC.query">查询</el-button>
                </el-col>
            </el-row>
            <divSelectTable :queryObj="editDLC.model" :queryFunc="querySuppliers" :selectObj="edit.model.suppliers"
                ref="supplierSelect"></divSelectTable>
            <el-row>
                <el-col :span="24">
                    <h3>已选择：</h3>
                </el-col>
                <el-col :span="22" :offset="2" v-for="supplier in editDLC.model.suppliers" :key="supplier.id">
                    {{ supplier.name }}
                </el-col>
            </el-row>
        </el-dialog>

        <el-dialog v-model="edit.moneyDialogVisivle" title="产品价格信息编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="150px" :rules="rules" ref="editForm">
                <el-form-item label="名称">
                    <el-input v-model.trim="edit.model.name" disabled />
                </el-form-item>
                <el-form-item label="采购价格(元)" prop="attribute.purchasePrice">
                    <el-input-number v-model="edit.model.attribute.purchasePrice" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
                <el-form-item label="采购价格(美元)" prop="attribute.purchasePriceUSD">
                    <el-input-number v-model="edit.model.attribute.purchasePriceUSD" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
                <el-form-item label="标准价格(元)" prop="attribute.standardPrice">
                    <el-input-number v-model="edit.model.attribute.standardPrice" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
                <el-form-item label="标准价格(美元)" prop="attribute.standardPriceUSD">
                    <el-input-number v-model="edit.model.attribute.standardPriceUSD" :controls="false" :min="0"
                        :max="9999999999" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="edit.submitMoney"
                            :disabled="edit.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { addProduct, editProductBase, editProductAttribute, queryProduct, queryProducts } from "@/api/product"
import { querySuppliers } from "@/api/supplier"
import { queryAllProductType } from "@/api/product_type"
import { message, messageForCRUD } from '@/components/divMessage/index'
import { reg_money, reg_number } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'
import divSelectTable from '@/components/divSelectTable/index.vue'

const addForm = ref(null)
const supplierSelect = ref(null)
const editForm = ref(null)
const rules = reactive({
    typeID: [
        { required: true, message: '请选择产品类型', trigger: 'blur' },
    ],
    name: [
        { required: true, message: '请输入产品名称', trigger: 'blur' },
    ],
    attribute: {
        purchasePrice: [
            { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
        ],
        purchasePriceUSD: [
            { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
        ],
        standardPrice: [
            { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
        ],
        standardPriceUSD: [
            { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
        ],
    },
    numberCount: [
        { required: true, pattern: reg_number, message: '请输入有效数字', trigger: 'blur' }
    ],
    unit: [
        { required: true, message: '请输入库存单位', trigger: 'blur' },
    ],
})

const base = reactive({
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
                width: "5%",
            },
            {
                prop: "name",
                label: "名称",
                width: "10%",
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
                prop: "numberCount",
                label: "库存数量",
                width: "8%",
            },
            {
                prop: "number",
                label: "可售数量",
                width: "8%",
            },
            {
                prop: "unit",
                label: "单位",
                width: "5%",
            },
            {
                prop: "attribute.standardPrice",
                label: "标准售价(元)",
                width: "8%",
            },
            {
                prop: "attribute.standardPriceUSD",
                label: "标准售价(美元)",
                width: "8%",
            },
            {
                type: "boolean",
                prop: "isFree",
                label: "小零配件",
                width: "8%",
            },
            {
                type: "operation",
                label: "操作",
                width: "25%",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "查看",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openViewDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "基础编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditBaseDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "价格编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditMoneyDialog(index, row)
                    },
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
        queryProducts(base.model, base.pageData).then((res) => {
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
        add.dialogVisible = true
    },
    openAddDLCDialog: () => {
        addDLC.dialogVisible = true
    },
    openViewDialog: (index, row) => {
        queryProduct(row.id).then((res) => {
            if (res.status == 1) {
                view.model = res.data
            }
        })
        view.dialogVisible = true
    },
    openEditBaseDialog: (index, row) => {
        queryProduct(row.id).then((res) => {
            if (res.status == 1) {
                edit.model = res.data
            }
        })
        edit.baseDialogVisible = true
    },
    openEditDLCDialog: (index, row) => {
        editDLC.dialogVisible = true
    },
    openEditMoneyDialog: (index, row) => {
        queryProduct(row.id).then((res) => {
            if (res.status == 1) {
                edit.model = res.data
            }
        })
        edit.moneyDialogVisivle = true
    },
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        name: "",
        brand: "",
        specification: "",
        number: 0,
        numberCount: 0,
        unit: "",
        deliveryCycle: "",
        remark: "",
        typeID: null,
        attribute: {
            purchasePrice: 0,
            purchasePriceUSD: 0,
            standardPrice: 0,
            standardPriceUSD: 0,
        },
        suppliers: [],
        isFree: false,
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addProduct(add.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add.model.name, "添加成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(add.model.name, "添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        name: "",
                        brand: "",
                        specification: "",
                        number: 0,
                        numberCount: 0,
                        unit: "",
                        deliveryCycle: "",
                        remark: "",
                        typeID: null,
                        attribute: {
                            purchasePrice: 0,
                            purchasePriceUSD: 0,
                            standardPrice: 0,
                            standardPriceUSD: 0,
                        },
                        suppliers: []
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    }
})

const addDLC = reactive({
    dialogVisible: false,
    model: {
        name: "",
        linkman: "",
        phone: "",
    },
    query: () => {
        supplierSelect.value.base.query()
    }
})

const view = reactive({
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
        numberCount: 0,
        unit: "",
        deliveryCycle: "",
        remark: "",
        isFree: false,
    }
})

const edit = reactive({
    baseDialogVisible: false,
    moneyDialogVisivle: false,
    submitDisabled: false,
    model: {
        id: null,
        typeID: null,
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
        numberCount: 0,
        unit: "",
        deliveryCycle: "",
        remark: "",
        isFree: false
    },
    submitBase: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editProductBase(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.baseDialogVisible = false
                    edit.model = {
                        id: null,
                        typeID: null,
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
                        numberCount: 0,
                        unit: "",
                        deliveryCycle: "",
                        remark: "",
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    },
    submitMoney: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editProductAttribute(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.moneyDialogVisivle = false
                    edit.model = {
                        id: null,
                        typeID: null,
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
                        numberCount: 0,
                        unit: "",
                        deliveryCycle: "",
                        remark: "",
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    },
})

const editDLC = reactive({
    dialogVisible: false,
    model: {
        name: "",
        linkman: "",
        phone: "",
    },
    query: () => {
        supplierSelect.value.base.query()
    }
})

onBeforeMount(() => {
    base.query()
    queryAllProductType().then((res) => {
        if (res.status == 1) {
            base.productTypes = res.data
        }
    })
})
</script>
