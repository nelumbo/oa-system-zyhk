<template>
    <el-row :gutter="20">
        <el-col :span="5" :offset="1">
            <el-input v-model="base.model.no" placeholder="合同编号" clearable maxlength="50" />
        </el-col>
        <el-col :span="5">
            <el-select v-model="base.model.regionID" placeholder="省份" clearable style="width: 100%;">
                <el-option v-for="item in base.regions" :key="item.id" :label="item.name" :value="item.id" />
            </el-select>
        </el-col>
        <el-col :span="5">
            <el-input v-model="base.model.customer.customerCompany.name" placeholder="客户单位" clearable maxlength="25" />
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
            <el-button type="success" @click="base.jumpEntryPage()">录入</el-button>
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

    <el-dialog v-model="view.dialogVisible" title="合同查看" width="75%" :show-close="false">
        <el-divider content-position="left">
            <h2>基本信息</h2>
        </el-divider>
        <el-form :model="view.model" label-width="120px">
            <el-row>
                <el-col :span="6">
                    <el-form-item label="合同编号">
                        <el-input v-model="view.model.no" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="录入时间">
                        <el-input v-model="view.model.createDate" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="办事处">
                        <el-input v-model="view.model.office.name" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="业务员">
                        <el-input v-model="view.model.employee.name" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="区域">
                        <el-input v-model="view.model.region.name" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="客户公司">
                        <el-input v-model="view.model.customer.customerCompany.name" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="客户">
                        <el-input v-model="view.model.customer.name" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6"></el-col>
                <el-col :span="6">
                    <el-form-item label="签订日期">
                        <el-input v-model="view.model.contractDate" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="交货日期">
                        <el-input v-model="view.model.estimatedDeliveryDate" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="实际交货日期">
                        <el-input v-model="view.model.endDeliveryDate" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="最后回款日期">
                        <el-input v-model="view.model.endPaymentDate" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="特殊合同">
                        <el-input v-model="view.isSpecialString" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="预存款合同">
                        <el-input v-model="view.isPreDepositString" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="12"></el-col>
                <el-col :span="6">
                    <el-form-item label="签订单位">
                        <el-input v-model="view.model.vendor.name" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="付款类型">
                        <el-input v-model="view.payTypeString" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="总金额">
                        <el-input v-model="view.model.totalAmount" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="回款额">
                        <el-input v-model="view.model.paymentTotalAmount" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="发票类型">
                        <el-input v-model="view.invoiceTypeString" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="18"></el-col>
                <el-col :span="24">
                    <el-form-item label="发票内容">
                        <el-input v-model="view.model.invoiceContent" type="textarea"
                            :autosize="{ minRows: 1, maxRows: 9 }" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="24">
                    <el-form-item label="付款方式">
                        <el-input v-model="view.model.paymentContent" type="textarea"
                            :autosize="{ minRows: 1, maxRows: 9 }" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="24">
                    <el-form-item label="备注">
                        <el-input v-model="view.model.remark" type="textarea" :autosize="{ minRows: 1, maxRows: 9 }"
                            readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="合同状态">
                        <el-input v-model="view.statusString" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="生产状态">
                        <el-input v-model="view.productionStatusString" readonly />
                    </el-form-item>
                </el-col>
                <el-col :span="6">
                    <el-form-item label="回款状态">
                        <el-input v-model="view.collectionStatusString" readonly />
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
        <el-divider content-position="left" style="margin-top: 15px;">
            <h2>产品详情</h2>
        </el-divider>
        <divTable :columnObj="view.columnT" :tableData="view.model.tasks" :allShow="true" />
        <el-divider content-position="left" style="margin-top: 15px;">
            <h2>开票记录</h2>
        </el-divider>
        <divTable :columnObj="view.columnI" :tableData="view.model.invoices" :allShow="true" />
        <el-divider content-position="left" style="margin-top: 15px;">
            <h2>回款记录</h2>
        </el-divider>
        <divTable :columnObj="view.columnPP" :tableData="view.model.payments" :allShow="true"
            v-if="view.model.isPreDeposit" />
        <divTable :columnObj="view.columnP" :tableData="view.model.payments" :allShow="true" v-else />
        <el-divider content-position="left" style="margin-top: 15px;" v-if="view.model.isPreDeposit">
            <h2>提成记录</h2>
        </el-divider>
        <divTable :columnObj="view.columnAuto" :tableData="view.model.paymentAutos" :allShow="true"
            v-if="view.model.isPreDeposit" />
    </el-dialog>

    <el-dialog v-model="viewDLC.dialogVisible" title="物流备注" width="50%" :show-close="false">
        <el-form :model="viewDLC.model" label-width="120px">
            <el-form-item label="物流备注">
                <el-input v-model="viewDLC.model.shipmentRemark" readonly />
            </el-form-item>
        </el-form>
    </el-dialog>

    <el-dialog v-model="edit.dialogVisible" title="编辑" width="75%" :show-close="false">
        <divTable :columnObj="edit.column" :tableData="edit.pageData" :allShow="true" />
    </el-dialog>

    <el-dialog v-model="editDLC.dialogVisible" title="删除任务" width="50%" :show-close="false">
        <h1>是否确定删除该条任务？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="editDLC.submit" :disabled="editDLC.submitDisabled">确定</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="add.dialogVisible" title="采购" width="75%" :show-close="false">
        <el-form :model="viewDLC.model" label-width="120px">
            <el-row>
                <el-col :span="12">
                    <el-form-item label="合同编号">
                        <el-input v-model="add.contract.no" disabled />
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="可用预存款">
                        <el-input v-model="add.contract.preDeposit" disabled />
                    </el-form-item>
                </el-col>
            </el-row>
        </el-form>
        <el-divider />
        <el-form :model="add.model" label-width="100px">
            <el-row :gutter="20">
                <el-col :span="7" :offset="1">
                    <el-input v-model="add.model.name" placeholder="名称" clearable maxlength="25" />
                </el-col>
                <el-col :span="7">
                    <el-input v-model="add.model.version" placeholder="型号" clearable maxlength="25" />
                </el-col>
                <el-col :span="7">
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

    <el-dialog v-model="addDLC.dialogVisible" :title="addDLC.product.name" width="50%" :show-close="false">
        <el-form :model="addDLC.model" label-width="60px" :rules="rules" ref="addForm">
            <el-form-item label="数量" prop="number">
                <el-input-number v-model="addDLC.model.number" :controls="false" :min="1" />
            </el-form-item>
            <el-form-item label="价格" prop="price">
                <el-input-number v-model="addDLC.model.price" :controls="false" :min="0" />
            </el-form-item>
            <el-form-item label="备注">
                <el-input v-model="addDLC.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
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

    <el-dialog v-model="del.dialogVisible" title="合同删除" width="50%" :show-close="false">
        <h1>是否确定删除该合同？</h1>
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
import { useRouter } from 'vue-router'
import { ref, reactive, onBeforeMount, computed } from 'vue'
import {
    contractStatusItems, productionStatusItems, collectionStatusItems,
    payTypeItems, invoiceTypeItems, taskStatusItems, boolItems,
} from '@/utils/magic'
import { queryAllRegion } from "@/api/region"
import { queryMyContracts } from "@/api/my"
import { delContract, queryContract } from "@/api/contract"
import { queryProduct, queryProducts } from "@/api/product"
import { addTask, delTask } from "@/api/contract_task"
import { message } from '@/components/divMessage/index'
import { reg_money, reg_number } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const router = useRouter()
const addForm = ref(null)
const rules = reactive({
    number: [
        { required: true, pattern: reg_number, message: '请输入大于零的有效数字', trigger: 'blur' },
    ],
    price: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
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
        status: 2,
        productionStatus: null,
        collectionStatus: null,
        payType: null,
        startDate: "",
        endDate: "",
        isSpecialNum: null,
        isPreDepositNum: null,
    },
    column: {
        headers: [
            {
                prop: "no",
                label: "合同编号",
                width: "15%",
            },
            {
                prop: "customer.customerCompany.name",
                label: "客户单位",
                width: "13%",
            },
            {
                prop: "customer.name",
                label: "客户",
                width: "5%",
            },
            {
                prop: "estimatedDeliveryDate",
                label: "交货日期",
                width: "8%",
            },
            {
                prop: "endDeliveryDate",
                label: "实际交货日期",
                width: "8%",
            },
            {
                type: "boolean",
                prop: "isPreDeposit",
                label: "预存款",
                width: "5%",
            },
            {
                type: "contractPreDeposit",
                prop: "preDeposit",
                label: "剩余预存款",
                width: "8%",
            },
            {
                prop: "totalAmount",
                label: "总金额",
                width: "5%",
            },
            {
                type: "contractNotPayment",
                prop: "notPaymentTotalAmount",
                label: "未回款",
                width: "8%",
            },
            {
                type: "contractType",
                prop: "status",
                label: "状态",
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
                        label: "查看",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openViewDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 2 && row.isPreDeposit) {
                                return true
                            }
                            return false
                        },
                        label: "采购",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openAddDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.isPreDeposit) {
                                return true
                            }
                            return false
                        },
                        label: "编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openEditDialog(index, row)
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
                    },
                ]
            },
        ],
        cellStyle: ({ row, column, rowIndex, columnIndex }) => {
            if (columnIndex == 0) {
                if (row.endDeliveryDate != "") {
                    var old_date = new Date(row.endDeliveryDate)
                    if (row.endPaymentDate != "") {
                        var new_date = new Date(row.endPaymentDate)
                    } else {
                        var new_date = new Date()
                    }
                    var difftime = (new_date - old_date) / 1000;
                    if (difftime > (-7 * 24 * 60 * 60) && difftime <= (60 * 24 * 60 * 60)) {
                        return {
                            backgroundColor: '#FF8C00',
                            color: '#000',
                        }
                    } else if (difftime > (60 * 24 * 60 * 60)) {
                        return {
                            backgroundColor: '#FF4500',
                            color: '#000',
                        }
                    }
                }
            } else if (columnIndex == 3) {
                var old_date = new Date(row.estimatedDeliveryDate)
                if (row.endDeliveryDate != "") {
                    var new_date = new Date(row.endDeliveryDate)
                } else {
                    var new_date = new Date()
                }
                var difftime = (new_date - old_date) / 1000
                if (row.no == "Bjscistar20230102-AAAZBC003") {
                    console.log(difftime)
                }
                if (difftime >= 1 * 24 * 60 * 60) {
                    return {
                        backgroundColor: '#FF4500',
                        color: '#000'
                    }
                }
            }
        }
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
    jumpEntryPage: () => {
        router.push("entry")
    },
    openViewDialog: (index, row) => {
        queryContract(row.id).then((res) => {
            if (res.status == 1) {
                if (res.data.isPreDeposit) {
                    let arrs = units.paymentsSeparation(res.data.payments)
                    res.data.payments = arrs[0]
                    res.data.paymentAutos = arrs[1]
                }
                view.model = res.data
            }
        })
        view.dialogVisible = true
    },
    openAddDialog: (index, row) => {
        add.contract = row
        add.query()
        add.dialogVisible = true
    },
    openEditDialog: (index, row) => {
        queryContract(row.id).then((res) => {
            if (res.status == 1) {
                edit.model.contractID = row.id
                edit.pageData = res.data.tasks
            }
        })
        edit.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.dialogVisible = true
    }
})

const view = reactive({
    dialogVisible: false,
    isSpecialString: computed(() => {
        if (view.model.isSpecial) {
            return "是"
        } else {
            return "否"
        }
    }),
    isPreDepositString: computed(() => {
        if (view.model.isPreDeposit) {
            return "是"
        } else {
            return "否"
        }
    }),
    payTypeString: computed(() => {
        var temp = "";
        payTypeItems.some((item) => {
            if (item.value == view.model.payType) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    invoiceTypeString: computed(() => {
        var temp = "";
        invoiceTypeItems.some((item) => {
            if (item.value == view.model.invoiceType) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    statusString: computed(() => {
        var temp = "";
        contractStatusItems.some((item) => {
            if (item.value == view.model.status) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    productionStatusString: computed(() => {
        var temp = "";
        productionStatusItems.some((item) => {
            if (item.value == view.model.productionStatus) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    collectionStatusString: computed(() => {
        var temp = "";
        collectionStatusItems.some((item) => {
            if (item.value == view.model.collectionStatus) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    model: {
        no: "",
        createDate: "",
        office: {
            name: "",
        },
        employee: {
            name: "",
        },
        region: {
            name: "",
        },
        customer: {
            customerCompany: {
                name: "",
            },
            name: "",
        },
        contractDate: "",
        estimatedDeliveryDate: "",
        endDeliveryDate: "",
        endPaymentDate: "",
        isSpecial: false,
        isPreDeposit: false,
        vendor: {
            name: "",
        },
        payType: null,
        totalAmount: 0,
        paymentTotalAmount: 0,
        invoiceType: null,
        invoiceContent: "",
        paymentContent: "",
        remark: "",
        status: null,
        productionStatus: null,
        collectionStatus: null,
        tasks: [],
        invoices: [],
        payments: [],
        paymentAutos: [],
    },
    columnT: {
        headers: [
            {
                prop: "id",
                label: "ID",
                width: "5%",
            },
            {
                prop: "product.type.name",
                label: "产品类型",
                width: "5%",
            },
            {
                prop: "product.name",
                label: "产品名称",
                width: "10%",
            },
            {
                prop: "number",
                label: "数量",
                width: "5%",
            },
            {
                prop: "product.unit",
                label: "单位",
                width: "5%",
            },
            {
                prop: "productAttribute.standardPrice",
                label: "标准定价(人民币)",
                width: "10%",
            },
            {
                prop: "productAttribute.standardPriceUSD",
                label: "标准定价(美元)",
                width: "10%",
            },
            {
                prop: "price",
                label: "售卖单价",
                width: "10%",
            },
            {
                prop: "totalPrice",
                label: "售卖总价",
                width: "10%",
            },
            {
                type: "employees",
                prop: "employees",
                label: "负责人",
                width: "10%",
            },
            {
                type: "transform",
                prop: "status",
                label: "状态",
                items: taskStatusItems,
                width: "10%",
            },
            {
                type: "operation",
                label: "操作",
                width: "10%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == 6) {
                                return true
                            }
                            return false
                        },
                        label: "查看快递单号",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => view.openViewDLCDialog(index, row)
                    },
                ]
            },
        ]
    },
    columnI: {
        headers: [
            {
                prop: "createDate",
                label: "创建时间",
            },
            {
                prop: "employee.name",
                label: "创建人",
            },
            {
                prop: "no",
                label: "发票号",
            },
            {
                prop: "money",
                label: "金额",
            },
        ],
    },
    columnP: {
        headers: [
            {
                prop: "createDate",
                label: "创建时间",
            },
            {
                prop: "employee.name",
                label: "创建人",
            },
            {
                prop: "paymentDate",
                label: "回款时间",
            },
            {
                prop: "task.id",
                label: "任务id",
            },
            {
                prop: "task.product.name",
                label: "产品",
            },
            {
                prop: "task.product.type.name",
                label: "产品类型",
            },
            {
                prop: "money",
                label: "回款金额",
            },
            {
                prop: "theoreticalPushMoney",
                label: "理论提成",
            },
            {
                prop: "fine",
                label: "回款延迟扣除",
            },
            {
                prop: "pushMoney",
                label: "实际提成",
            },
            {
                prop: "businessMoney",
                label: "业务费用",
            },
        ],
    },
    columnPP: {
        headers: [
            {
                prop: "createDate",
                label: "创建时间",
            },
            {
                prop: "employee.name",
                label: "创建人",
            },
            {
                prop: "paymentDate",
                label: "回款时间",
            },
            {
                prop: "money",
                label: "回款金额",
            },
        ],
    },
    columnAuto: {
        headers: [
            {
                prop: "createDate",
                label: "创建时间",
            },
            {
                prop: "task.id",
                label: "任务id",
            },
            {
                prop: "task.product.name",
                label: "产品",
            },
            {
                prop: "task.product.type.name",
                label: "产品类型",
            },
            {
                prop: "money",
                label: "回款金额",
            },
            {
                prop: "theoreticalPushMoney",
                label: "理论提成",
            },
            {
                prop: "fine",
                label: "回款延迟扣除",
            },
            {
                prop: "pushMoney",
                label: "实际提成",
            },
            {
                prop: "businessMoney",
                label: "业务费用",
            },
        ],
    },
    openViewDLCDialog: (index, row) => {
        viewDLC.model.shipmentRemark = row.shipmentRemark
        viewDLC.dialogVisible = true
    }
})

const viewDLC = reactive({
    dialogVisible: false,
    model: {
        shipmentRemark: "",
    }
})

const edit = reactive({
    dialogVisible: false,
    model: {
        contractID: 0,
    },
    column: {
        headers: [
            {
                prop: "id",
                label: "ID",
                width: "5%",
            },
            {
                prop: "product.type.name",
                label: "产品类型",
                width: "5%",
            },
            {
                prop: "product.name",
                label: "产品名称",
                width: "10%",
            },
            {
                prop: "number",
                label: "数量",
                width: "5%",
            },
            {
                prop: "product.unit",
                label: "单位",
                width: "5%",
            },
            {
                prop: "productAttribute.standardPrice",
                label: "标准定价(人民币)",
                width: "10%",
            },
            {
                prop: "productAttribute.standardPriceUSD",
                label: "标准定价(美元)",
                width: "10%",
            },
            {
                prop: "price",
                label: "售卖单价",
                width: "10%",
            },
            {
                prop: "totalPrice",
                label: "售卖总价",
                width: "10%",
            },
            {
                type: "employees",
                prop: "employees",
                label: "负责人",
                width: "10%",
            },
            {
                type: "transform",
                prop: "status",
                label: "状态",
                items: taskStatusItems,
                width: "10%",
            },
            {
                type: "operation",
                label: "操作",
                width: "10%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == -1) {
                                return true
                            }
                            return false
                        },
                        label: "删除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => edit.openEditDLCDialog(index, row)
                    },
                ]
            },
        ]
    },
    pageData: [],
    openEditDLCDialog: (index, row) => {
        editDLC.model.id = row.id
        editDLC.dialogVisible = true
    },
})

const editDLC = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0
    },
    submit: () => {
        editDLC.submitDisabled = true
        delTask(editDLC.model.id).then((res) => {
            if (res.status == 1) {
                message("删除成功", "success")
                queryContract(edit.model.contractID).then((res) => {
                    if (res.status == 1) {
                        edit.pageData = res.data.tasks
                    }
                })
            } else {
                message("删除失败", "error")
            }
            editDLC.dialogVisible = false
            editDLC.model = {
                id: 0,
            }
            editDLC.submitDisabled = false
        })
    }
})

const add = reactive({
    dialogVisible: false,
    contract: {
        id: "",
        no: "",
        preDeposit: 0,
    },
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
                width: "15%",
            },
            {
                prop: "version",
                label: "型号",
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
    handleSizeChange: (e) => {
        add.pageData.pageSize = e
        add.pageData.pageNo = 1
        add.query()
    },
    handleCurrentChange: (e) => {
        add.pageData.pageNo = e
        add.query()
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
        addDLC.model.contractID = add.contract.id
        addDLC.model.productID = row.id
        addDLC.model.number = 1
        if (add.contract.payType == 1) {
            addDLC.model.price = row.attribute.standardPrice
        } else {
            addDLC.model.price = row.attribute.standardPriceUSD
        }
        addDLC.model.remark = ""
        addDLC.product.name = row.name
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
    product: {
        name: "",
    },
    model: {
        contractID: null,
        productID: null,
        number: 0,
        price: 0,
        totalPrice: 0,
        remark: "",
        product: {
            name: "",
        }
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                addDLC.submitDisabled = true
                addDLC.model.totalPrice = addDLC.model.number * addDLC.model.price
                addTask(addDLC.model).then((res) => {
                    if (res.status == 1) {
                        message("添加成功", "success")
                        add.contract.preDeposit -= addDLC.model.totalPrice
                        base.query()
                    } else {
                        message("添加失败", "error")
                    }
                    addDLC.dialogVisible = false
                    addDLC.product = {
                        name: "",
                    }
                    addDLC.model = {
                        contractID: null,
                        productID: null,
                        number: "",
                        price: "",
                        totalPrice: "",
                        remark: "",
                    }
                    addDLC.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    }
})

const del = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
    },
    submit: () => {
        del.submitDisabled = true
        delContract(del.model.id).then((res) => {
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

const units = reactive({
    paymentsSeparation: (payments) => {
        let arr1 = payments.filter(
            function (item) {
                if (item.task.id == 0) {
                    return item
                }
            })
        let arr2 = payments.filter(
            function (item) {
                if (item.task.id > 0) {
                    return item
                }
            })

        if (!arr1) {
            arr1 = []
        }
        if (!arr2) {
            arr2 = []
        }

        return [arr1, arr2]
    },
})

onBeforeMount(() => {
    queryAllRegion().then((res) => {
        if (res.status == 1) {
            base.regions = res.data
        }
    })
    base.query()
})
</script>