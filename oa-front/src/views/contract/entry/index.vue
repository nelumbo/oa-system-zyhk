<template>
    <div>
        <el-dialog v-model="pre.dialogVisible" title="暂存合同" width="75%" :close-on-click-modal="false">
            <divTable :columnObj="pre.column" :tableData="pre.tableData" :pageData="pre.pageData"
                :handleSizeChange="pre.handleSizeChange" :handleCurrentChange="pre.handleCurrentChange" />
            <el-row style="margin-top: 30px;">
                <el-col :span="2" :offset="11">
                    <el-button type="success" @click="pre.dialogVisible = false">新合同</el-button>
                </el-col>
            </el-row>
        </el-dialog>
        <el-dialog v-model="preDel.dialogVisible" title="暂存合同删除" width="50%" :show-close="false">
            <h1>是否确定删除该条暂存合同？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="preDel.submit" :disabled="preDel.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-divider content-position="left">
            <h2>基本信息</h2>
        </el-divider>
        <el-form :model="base.model" label-width="400px" :rules="rules" ref="baseForm" style="width:70%;">
            <el-form-item label="省份" style="width: 100%;" prop="customerID">
                <el-input v-model="dlc.customer.customerCompany.region.name" disabled style="width:100%;" />
            </el-form-item>
            <el-form-item label="客户单位" style="width: 100%;" prop="customerID">
                <el-input v-model="dlc.customer.customerCompany.name" disabled style="width:100%;" />
            </el-form-item>
            <el-form-item label="客户名称" style="width: 100%;" prop="customerID">
                <el-input v-model="dlc.customer.name" disabled style="width:85%;" />
                <di style="width: 1%;"></di>
                <el-button type="primary" @click="base.openDLCDialog" style="width: 14%;">选择客户</el-button>
            </el-form-item>
            <el-form-item label="客户ID" style="width: 100%;" prop="customerID" v-show="false">
                <el-input v-model="base.model.customerID" disabled style="width:100%;" />
            </el-form-item>
            <el-form-item label="签订单位" prop="vendorID">
                <el-select v-model="base.model.vendorID" style="width: 100%;">
                    <el-option v-for="item in base.vendors" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-form-item>
            <el-form-item label="签订日期" prop="contractDate">
                <el-date-picker v-model="base.model.contractDate" type="date" style="width: 100%;" />
            </el-form-item>
            <el-form-item label="交货日期" prop="estimatedDeliveryDate">
                <el-date-picker v-model="base.model.estimatedDeliveryDate" type="date" style="width: 100%;" />
            </el-form-item>
            <el-form-item label="付款类型" prop="payType">
                <el-radio-group v-model="base.model.payType" @change="base.changePayType">
                    <el-radio v-for="item in payTypeItems" :label="item.value">{{ item.label }}</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="预存款" prop="isPreDeposit" @change="base.changeIsPreDeposit">
                <el-radio-group v-model="base.model.isPreDeposit">
                    <el-radio v-for="item in isPreDepositItems" :label="item.value">{{ item.label }}</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="预存款金额" prop="preDepositRecord" v-if="base.model.isPreDeposit">
                <el-input-number v-model="base.model.preDepositRecord" :controls="false" :min="0" :max="9999999999" />
            </el-form-item>
            <el-form-item label="特殊合同" prop="isSpecial">
                <el-radio-group v-model="base.model.isSpecial">
                    <el-radio v-for="item in isSpecialItems" :label="item.value">{{ item.label }}</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="开票类型" prop="invoiceType">
                <el-radio-group v-model="base.model.invoiceType">
                    <el-radio v-for="item in invoiceTypeItems" :label="item.value">{{ item.label }}</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="开票内容和要求" prop="invoiceContent">
                <el-input v-model="base.model.invoiceContent" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="500" v-if="base.model.invoiceType != 1" />
            </el-form-item>
            <el-form-item label="付款方式" prop="paymentContent">
                <el-input v-model="base.model.paymentContent" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="500" />
            </el-form-item>
            <el-form-item label="备注(发货地址等，如是经销商，须说明终端客户)" prop="remark">
                <el-input v-model="base.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
            </el-form-item>
            <el-form-item label="合同总金额" style="width: 100%;">
                <el-input v-model="base.model.totalPrice" disabled style="width:100%;" />
            </el-form-item>
        </el-form>

        <el-dialog v-model="dlc.dialogVisible" title="客户选择" width="50%" :show-close="false">
            <el-row :gutter="20">
                <el-col :span="5" :offset="1">
                    <el-select v-model="dlc.model.customerCompany.regionID" placeholder="省份" clearable style="width: 100%;">
                        <el-option v-for="item in dlc.regions" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                </el-col>
                <el-col :span="5">
                    <el-input v-model="dlc.model.customerCompany.name" placeholder="公司" clearable maxlength="25" />
                </el-col>
                <el-col :span="5">
                    <el-input v-model="dlc.model.name" placeholder="姓名" clearable maxlength="25" />
                </el-col>
                <el-col :span="5">
                    <el-input v-model="dlc.model.researchGroup" placeholder="课题组" clearable maxlength="25" />
                </el-col>
                <el-col :span="1">
                    <el-button type="primary" @click="dlc.query">查询</el-button>
                </el-col>
            </el-row>
            <divSelectTable :queryObj="dlc.model" :queryFunc="queryCustomers" :chooseFunc="dlc.choose" :isRadio="true"
                :headers="dlc.headers" ref="customerSelect"></divSelectTable>
        </el-dialog>
    </div>
    <div style="margin-top: 30px;" v-if="!base.model.isPreDeposit">
        <el-divider content-position="left">
            <h2>产品库</h2>
        </el-divider>
        <el-row :gutter="20">
            <el-col :span="6" :offset="2">
                <el-select v-model="bank.model.typeID" placeholder="类型" clearable style="width: 100%;">
                    <el-option v-for="item in bank.productTypes" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="6">
                <el-input v-model="bank.model.name" placeholder="名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="6">
                <el-input v-model="bank.model.specification" placeholder="规格" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="bank.query">查询</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="bank.column" :tableData="bank.tableData" :pageData="bank.pageData"
            :handleSizeChange="bank.handleSizeChange" :handleCurrentChange="bank.handleCurrentChange" />

        <el-dialog v-model="bankView.dialogVisible" title="产品查看" width="50%" :show-close="false">
            <el-form :model="bankView.model" label-width="150px">
                <el-form-item label="类型">
                    <el-input v-model.trim="bankView.model.type.name" disabled />
                </el-form-item>
                <el-form-item label="名称">
                    <el-input v-model.trim="bankView.model.name" disabled />
                </el-form-item>
                <el-form-item label="品牌">
                    <el-input v-model.trim="bankView.model.brand" disabled />
                </el-form-item>
                <el-form-item label="规格">
                    <el-input v-model.trim="bankView.model.specification" disabled />
                </el-form-item>
                <el-form-item label="供应商">
                    <el-row>
                        <el-col :span="24" v-for="supplier in bankView.model.suppliers" :key="supplier.id">
                            {{ supplier.name }}
                        </el-col>
                    </el-row>
                </el-form-item>
                <el-form-item label="采购价格(元)">
                    <el-input v-model.trim="bankView.model.attribute.purchasePrice" disabled />
                </el-form-item>
                <el-form-item label="采购价格(美元)">
                    <el-input v-model.trim="bankView.model.attribute.purchasePriceUSD" disabled />
                </el-form-item>
                <el-form-item label="标准售价(元)">
                    <el-input v-model.trim="bankView.model.attribute.standardPrice" disabled />
                </el-form-item>
                <el-form-item label="标准售价(美元)">
                    <el-input v-model.trim="bankView.model.attribute.standardPriceUSD" disabled />
                </el-form-item>
                <el-form-item label="库存数量" prop="number">
                    <el-input v-model.trim="bankView.model.number" disabled />
                </el-form-item>
                <el-form-item label="库存单位" prop="unit">
                    <el-input v-model.trim="bankView.model.unit" disabled />
                </el-form-item>
                <el-form-item label="供货周期">
                    <el-input v-model.trim="bankView.model.deliveryCycle" disabled />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="bankView.model.remark" type="textarea" :autosize=true disabled />
                </el-form-item>
            </el-form>
        </el-dialog>

        <el-dialog v-model="cartAdd.dialogVisible" :title="cartAdd.model.product.name" width="50%" :show-close="false">
            <el-form :model="cartAdd.model" label-width="60px" :rules="rules" ref="cartAddForm">
                <el-form-item label="数量" prop="number">
                    <el-input-number v-model="cartAdd.model.number" :controls="false" :min="1" :max="99999" />
                </el-form-item>
                <el-form-item label="价格" prop="price">
                    <el-input-number v-model="cartAdd.model.price" :controls="false" :min="0" :max="9999999999" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="cartAdd.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="cartAdd.submit">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
        <el-divider content-position="left" style="margin-top: 15px;">
            <h2>需求清单</h2>
        </el-divider>
        <divTable :columnObj="cart.column" :tableData="cart.tasks" :allShow="true" />

        <el-dialog v-model="cartDel.dialogVisible" width="50%" :show-close="false">
            <h1>是否确定删除【{{ cartDel.model.name }}】？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="cartDel.submit">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
    <el-row style="margin-top: 30px;">
        <el-col :span="2" :offset="7">
            <el-button type="success" @click="base.save" :disabled="base.submitDisabled" size="large">暂存</el-button>
        </el-col>
        <el-col :span="2" :offset="6">
            <el-button type="primary" @click="base.submit" :disabled="base.submitDisabled" size="large">提交</el-button>
        </el-col>
    </el-row>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { ref, reactive, onBeforeMount, computed } from 'vue'
import { payTypeItems, isPreDepositItems, isSpecialItems, invoiceTypeItems } from '@/utils/magic'
import { queryMySaveContracts } from "@/api/my"
import { saveContract, addContract, delContract } from "@/api/contract"
import { queryAllVendor } from "@/api/vendor"
import { queryAllRegion } from "@/api/region"
import { queryAllProductType } from "@/api/product_type"
import { queryCustomers } from "@/api/customer"
import { queryProduct, queryProducts } from "@/api/product"
import { message } from '@/components/divMessage/index'
import { reg_number, reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'
import divSelectTable from '@/components/divSelectTable/index.vue'

const router = useRouter()
const customerSelect = ref(null)

const baseForm = ref(null)
const cartAddForm = ref(null)
const rules = reactive({
    customerID: [
        { required: true, message: '请选择客户', trigger: 'blur' },
    ],
    vendorID: [
        { required: true, message: '请选择签订单位', trigger: 'blur' },
    ],
    contractDate: [
        { required: true, message: '请选择签订日期', trigger: 'blur' },
    ],
    estimatedDeliveryDate: [
        { required: true, message: '请选择交货日期', trigger: 'blur' },
    ],
    payType: [
        { required: true, message: '请选择付款类型', trigger: 'blur' },
    ],
    isPreDeposit: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    preDepositRecord: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    isSpecial: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    invoiceType: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    price: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    number: [
        { required: true, pattern: reg_number, message: '请输入有效数字', trigger: 'blur' }
    ],
})

const pre = reactive({
    dialogVisible: false,
    column: {
        headers: [
            {
                prop: "createDate",
                label: "创建时间",
                width: "15%",
            },
            {
                prop: "customer.customerCompany.name",
                label: "客户单位",
                width: "10%",
            },
            {
                prop: "customer.name",
                label: "客户",
                width: "10%",
            },
            {
                prop: "contractDate",
                label: "签订日期",
                width: "15%",
            },
            {
                prop: "estimatedDeliveryDate",
                label: "交货日期",
                width: "15%",
            },
            {
                type: "boolean",
                prop: "isSpecial",
                label: "特殊合同",
                width: "10%",
            },
            {
                type: "boolean",
                prop: "isPreDeposit",
                label: "预存款合同",
                width: "10%",
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
                        label: "续填",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => pre.next(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "删除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => pre.openDelDialog(index, row)
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
        queryMySaveContracts(null, pre.pageData).then((res) => {
            if (res.status == 1) {
                pre.tableData = res.data.data
                pre.pageData.total = res.data.total
                pre.pageData.pageSize = res.data.pageSize
                pre.pageData.pageNo = res.data.pageNo
                if (pre.pageData.total > 0) {
                    pre.dialogVisible = true
                }
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
    next: (index, row) => {
        base.model.id = row.id
        dlc.customer.customerCompany.region.name = row.customer.customerCompany.region.name
        dlc.customer.customerCompany.name = row.customer.customerCompany.name
        dlc.customer.name = row.customer.name

        base.model.customerID = row.customerID
        base.model.regionID = row.regionID
        base.model.vendorID = row.vendorID
        base.model.contractDate = new Date(row.contractDate.replace(/-/g, "/"));
        base.model.estimatedDeliveryDate = new Date(row.estimatedDeliveryDate.replace(/-/g, "/"));
        base.model.payType = row.payType
        base.model.isSpecial = row.isSpecial
        base.model.isPreDeposit = row.isPreDeposit
        base.model.preDepositRecord = row.preDepositRecord
        base.model.invoiceType = row.invoiceType
        base.model.invoiceContent = row.invoiceContent
        base.model.paymentContent = row.paymentContent
        base.model.remark = row.remark
        base.model.tasks = []

        if (row.vendorID == 0) {
            base.model.vendorID = null
        }
        pre.dialogVisible = false
    },
    openDelDialog: (index, row) => {
        preDel.model.id = row.id
        preDel.dialogVisible = true
    }
})

const preDel = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
    },
    submit: () => {
        preDel.submitDisabled = true
        delContract(preDel.model.id).then((res) => {
            if (res.status == 1) {
                message("删除成功", "success")
                pre.query()
            } else {
                message("删除失败", "error")
            }
            preDel.dialogVisible = false
            preDel.model = {
                id: 0,
            }
            preDel.submitDisabled = false
        })
    }
})

const base = reactive({
    vendors: [],
    submitDisabled: false,
    model: {
        id: null,
        customerID: null,
        regionID: null,
        vendorID: null,
        contractDate: new Date(
            Date.now() - new Date().getTimezoneOffset() * 60000
        )
            .toISOString(),
        estimatedDeliveryDate: "",
        payType: 1,
        totalPrice: computed(() => {
            if (cart.tasks) {
                let temp = 0
                cart.tasks.forEach((item, index) => {
                    temp += item.totalPrice
                })
                return temp
            } else {
                return 0
            }
        }),
        isSpecial: false,
        isPreDeposit: false,
        preDepositRecord: 0,
        invoiceType: 2,
        invoiceContent: "",
        paymentContent: "",
        remark: "",
        tasks: [],
    },
    changePayType: () => {
        cart.tasks = []
    },
    changeIsPreDeposit: () => {
        cart.tasks = []
        base.model.preDepositRecord = 0
    },
    save: () => {
        base.submitDisabled = true
        saveContract(base.model).then((res) => {
            if (res.status == 1) {
                message("保存成功", "success")
                router.push("home")
            } else {
                message("保存失败", "error")
            }
            base.dialogVisible = false
            base.model = {
                regionID: null,
                customerID: null,
                vendorID: null,
                contractDate: new Date(
                    Date.now() - new Date().getTimezoneOffset() * 60000
                )
                    .toISOString(),
                estimatedDeliveryDate: "",
                payType: 1,
                totalPrice: computed(() => {
                    if (cart.tasks) {
                        let temp = 0
                        cart.tasks.forEach((item, index) => {
                            temp += item.totalPrice
                        })
                        return temp
                    } else {
                        return 0
                    }
                }),
                isSpecial: false,
                isPreDeposit: false,
                preDepositRecord: 0,
                invoiceType: 2,
                invoiceContent: "",
                paymentContent: "",
                remark: "",
                tasks: [],
            }
            base.submitDisabled = false
        })
    },
    submit: () => {
        baseForm.value.validate((valid) => {
            if (valid) {
                base.submitDisabled = true
                base.model.tasks = cart.tasks
                addContract(base.model).then((res) => {
                    if (res.status == 1) {
                        message("录入成功", "success")
                        router.push("home")
                    } else {
                        message("录入失败", "error")
                    }
                    base.dialogVisible = false
                    base.model = {
                        regionID: null,
                        customerID: null,
                        vendorID: null,
                        contractDate: new Date(
                            Date.now() - new Date().getTimezoneOffset() * 60000
                        )
                            .toISOString(),
                        estimatedDeliveryDate: "",
                        payType: 1,
                        totalPrice: computed(() => {
                            if (cart.tasks) {
                                let temp = 0
                                cart.tasks.forEach((item, index) => {
                                    temp += item.totalPrice
                                })
                                return temp
                            } else {
                                return 0
                            }
                        }),
                        isSpecial: false,
                        isPreDeposit: false,
                        preDepositRecord: 0,
                        invoiceType: 2,
                        invoiceContent: "",
                        paymentContent: "",
                        remark: "",
                        tasks: [],
                    }
                    base.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
    openDLCDialog: () => {
        queryAllRegion().then((res) => {
            if (res.status == 1) {
                dlc.regions = res.data
            }
        })
        dlc.dialogVisible = true
    },
})

const dlc = reactive({
    dialogVisible: false,
    regions: [],
    model: {
        name: "",
        researchGroup: "",
        customerCompany: {
            name: "",
        },
    },
    headers: [
        {
            prop: "customerCompany.region.name",
            label: "区域",
        },
        {
            prop: "customerCompany.name",
            label: "公司",
        },
        {
            prop: "name",
            label: "姓名",
        },
        {
            prop: "phone",
            label: "联系电话",
        },
    ],
    customer: {
        name: "",
        customerCompany: {
            name: "",
            region: {
                name: ""
            }
        },
    },
    choose: (index, row) => {
        base.model.regionID = row.customerCompany.regionID
        base.model.customerID = row.id
        dlc.customer.name = row.name
        dlc.customer.customerCompany.name = row.customerCompany.name
        dlc.customer.customerCompany.region.name = row.customerCompany.region.name
        dlc.dialogVisible = false
    },
    query: () => {
        customerSelect.value.base.query()
    }
})

const bank = reactive({
    productTypes: [],
    model: {
        typeID: null,
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
                width: "10%",
            },
            {
                prop: "brand",
                label: "品牌",
                width: "10%",
            },
            {
                prop: "specification",
                label: "规格",
                width: "10%",
            },
            {
                prop: "Version",
                label: "型号",
                width: "10%",
            },
            {
                prop: "number",
                label: "库存数量",
                width: "5%",
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
                fixed: "right",
                type: "operation",
                label: "操作",
                width: "300px",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "查看",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => bank.openBankViewDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "添加",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => bank.openCartAddDialog(index, row)
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
        queryProducts(bank.model, bank.pageData).then((res) => {
            if (res.status == 1) {
                bank.tableData = res.data.data
                bank.pageData.total = res.data.total
                bank.pageData.pageSize = res.data.pageSize
                bank.pageData.pageNo = res.data.pageNo
            } else {
                message("查询失败", "error")
            }
        })
    },
    handleSizeChange: (e) => {
        bank.pageData.pageSize = e
        bank.pageData.pageNo = 1
        bank.query()
    },
    handleCurrentChange: (e) => {
        bank.pageData.pageNo = e
        bank.query()
    },
    openBankViewDialog: (index, row) => {
        queryProduct(row.id).then((res) => {
            if (res.status == 1) {
                bankView.model = res.data
            }
        })
        bankView.dialogVisible = true
    },
    openCartAddDialog: (index, row) => {
        cartAdd.model.productID = row.id
        cartAdd.model.product.name = row.name
        cartAdd.model.number = 1
        cartAdd.model.unit = row.unit
        if (base.model.payType == 1) {
            cartAdd.model.price = row.attribute.standardPrice
        } else {
            cartAdd.model.price = row.attribute.standardPriceUSD
        }
        cartAdd.dialogVisible = true
    }
})

const bankView = reactive({
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

const cart = reactive({
    dialogVisible: false,
    tasks: [],
    column: {
        headers: [
            {
                prop: "product.name",
                label: "名称",
            },
            {
                prop: "number",
                label: "数量",
            },
            {
                prop: "unit",
                label: "单位",
            },
            {
                prop: "price",
                label: "单价",
            },
            {
                prop: "totalPrice",
                label: "总价",
            },
            {
                prop: "remark",
                label: "总价",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "删除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => cart.openCartDelDialog(index, row)
                    }
                ]
            },
        ]
    },
    openCartDelDialog: (index, row) => {
        cartDel.model.id = index
        cartDel.model.name = row.product.name
        cartDel.dialogVisible = true
    },
})

const cartAdd = reactive({
    dialogVisible: false,
    model: {
        productID: null,
        number: 1,
        unit: "",
        price: 0,
        totalPrice: 0,
        remark: "",
        product: {
            name: "",
        },
    },
    submit: () => {
        cartAddForm.value.validate((valid) => {
            if (valid) {
                cartAdd.model.totalPrice = cartAdd.model.number * cartAdd.model.price

                let isAdd = true
                cart.tasks.forEach((item, index) => {
                    if (item.productID == cartAdd.model.productID && item.price == cartAdd.model.price) {
                        cart.tasks[index].number = cartAdd.model.number
                        cart.tasks[index].price = cartAdd.model.price
                        cart.tasks[index].totalPrice = cartAdd.model.totalPrice
                        cart.tasks[index].remark = cartAdd.model.remark
                        isAdd = false
                        return
                    }
                })
                if (isAdd) {
                    cart.tasks.push(JSON.parse(JSON.stringify(cartAdd.model)))
                }

                cartAdd.dialogVisible = false
            } else {
                return false;
            }
        })

    }
})

const cartDel = reactive({
    dialogVisible: false,
    model: {
        id: null,
        name: "",
    },
    submit: () => {
        cart.tasks.splice(cartDel.model.id, 1)
        cartDel.dialogVisible = false
    }
})


onBeforeMount(() => {
    pre.query()
    queryAllVendor().then((res) => {
        if (res.status == 1) {
            base.vendors = res.data
        }
    })
    queryAllProductType().then((res) => {
        if (res.status == 1) {
            bank.productTypes = res.data
        }
    })
    bank.query()
})
</script>