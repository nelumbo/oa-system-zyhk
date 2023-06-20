<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-input v-model="base.model.no" placeholder="合同编号" clearable maxlength="50" />
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.employee.name" placeholder="业务员" clearable maxlength="25" />
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.collectionStatus" placeholder="回款状态" clearable style="width: 100%;">
                    <el-option v-for="item in collectionStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.invoiceType" placeholder="发票类型" clearable style="width: 100%;">
                    <el-option v-for="item in invoiceTypeItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-input v-model="base.model.customer.customerCompany.name" placeholder="客户单位" clearable maxlength="25" />
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.customer.name" placeholder="客户名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="5">
                <el-date-picker v-model="base.model.startDate" type="date" placeholder="开始时间" style="width: 100%;" />
            </el-col>
            <el-col :span="5">
                <el-date-picker v-model="base.model.endDate" type="date" placeholder="结束时间" style="width: 100%;" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-select v-model="base.model.isPreDepositNum" placeholder="预存款合同" clearable style="width: 100%;">
                    <el-option v-for="item in boolItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
        </el-row>

        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="view.dialogVisible" title="查看" width="75%" :show-close="false">
            <el-divider content-position="left">
                <h2>发票记录</h2>
            </el-divider>
            <divTable :columnObj="view.columnI" :tableData="view.model.invoices" :allShow="true" />
            <el-divider content-position="left" style="margin-top: 30px;">
                <h2>回款记录</h2>
            </el-divider>
            <divTable :columnObj="view.columnP" :tableData="view.model.payments" :allShow="true" />
        </el-dialog>

        <el-dialog v-model="viewP.dialogVisible" title="查看" width="75%" :show-close="false">
            <el-divider content-position="left">
                <h2>发票记录</h2>
            </el-divider>
            <divTable :columnObj="viewP.columnI" :tableData="viewP.model.invoices" :allShow="true" />
            <el-divider content-position="left" style="margin-top: 30px;">
                <h2>回款记录</h2>
            </el-divider>
            <divTable :columnObj="viewP.columnP" :tableData="viewP.model.payments" :allShow="true" />
            <el-divider content-position="left" style="margin-top: 30px;">
                <h2>提成计算</h2>
            </el-divider>
            <divTable :columnObj="viewP.columnAuto" :tableData="viewP.model.paymentAutos" :allShow="true" />
        </el-dialog>

        <el-dialog v-model="addI.dialogVisible" title="发票添加" width="75%" :show-close="false">
            <el-divider content-position="left">
                <h2>业务员备注</h2>
            </el-divider>
            <el-form :model="addI.model" label-width="100px">
                <el-form-item label="付款类型">
                    <el-input v-model="addI.payTypeString" readonly />
                </el-form-item>
                <el-form-item label="发票类型">
                    <el-input v-model="addI.invoiceTypeString" readonly />
                </el-form-item>
                <el-form-item label="合同总金额">
                    <el-input v-model="addI.contract.totalAmount" readonly />
                </el-form-item>
                <el-form-item label="发票内容">
                    <el-input v-model="addI.contract.invoiceContent" type="textarea" :autosize="{ minRows: 1, maxRows: 9 }"
                        readonly />
                </el-form-item>
            </el-form>
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>发票记录</h2>
            </el-divider>
            <divTable :columnObj="addI.column" :tableData="addI.contract.invoices" :allShow="true" />
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>添加发票</h2>
            </el-divider>
            <el-form :model="addI.model" label-width="80px" :rules="rules" ref="addIForm">
                <el-form-item label="发票号" prop="no">
                    <el-input v-model="addI.model.no" maxlength="100" />
                </el-form-item>
                <el-form-item label="金额" prop="money">
                    <el-input-number v-model="addI.model.money" :controls="false" :min="0" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="addI.submit" :disabled="addI.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="editI.dialogVisible" title="发票编辑" width="50%" :show-close="false">
            <el-form :model="editI.model" label-width="80px" :rules="rules" ref="editIForm">
                <el-form-item label="发票号" prop="no">
                    <el-input v-model="editI.model.no" maxlength="100" />
                </el-form-item>
                <el-form-item label="金额" prop="money">
                    <el-input-number v-model="editI.model.money" :controls="false" :min="0" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="editI.submit" :disabled="editI.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="delI.dialogVisible" title="发票删除" width="50%" :show-close="false">
            <h1>是否确定删除【{{ delI.model.no }}】？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="delI.submit" :disabled="delI.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="addP.dialogVisible" title="回款记录添加" width="75%" :show-close="false">
            <el-divider content-position="left">
                <h2>业务员备注</h2>
            </el-divider>
            <el-form :model="addP.model" label-width="100px">
                <el-form-item label="付款类型">
                    <el-input v-model="addP.payTypeString" readonly />
                </el-form-item>
                <el-form-item label="合同总金额">
                    <el-input v-model="addP.contract.totalAmount" readonly />
                </el-form-item>
                <el-form-item label="发票内容">
                    <el-input v-model="addP.contract.paymentContent" type="textarea" :autosize="{ minRows: 1, maxRows: 9 }"
                        readonly />
                </el-form-item>
            </el-form>
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>回款记录</h2>
            </el-divider>
            <divTable :columnObj="addP.columnP" :tableData="addP.contract.payments" :allShow="true"
                v-if="addP.contract.isPreDeposit" />
            <divTable :columnObj="addP.column" :tableData="addP.contract.payments" :allShow="true" v-else />
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>添加回款记录</h2>
            </el-divider>
            <el-form :model="addP.model" label-width="100px" :rules="rules" ref="addPForm">
                <div v-if="addP.contract.isPreDeposit">
                    <el-row>
                        <el-col :span="8">
                            <el-form-item label="金额" prop="money">
                                <el-input-number v-model="addP.model.money" :controls="false" :min="0" :max="9999999999"
                                    style="width: 100%;" />
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="8">
                            <el-form-item label="回款日期" prop="paymentDate">
                                <el-date-picker v-model="addP.model.paymentDate" type="date" style="width: 100%;" />
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
                <div v-else>
                    <el-row>
                        <el-col :span="8">
                            <el-form-item label="产品" prop="taskID">
                                <el-select v-model="addP.model.taskID" placeholder="请选择" @change="addP.taskChange"
                                    style="width: 100%;">
                                    <el-option v-for="item in addP.contract.tasks" :key="item.id"
                                        :label="item.product.name + '  [产品类型:' + item.product.type.name + ']  [任务id:' + item.id + ']'"
                                        :value="item.id" />
                                </el-select>
                            </el-form-item>
                        </el-col>
                        <el-form-item label="已回款金额" v-show="false" prop="taskID">
                            <el-input v-model="addP.task.taskID" disabled />
                        </el-form-item>
                        <el-col :span="8">
                            <el-form-item label="总金额">
                                <el-input v-model="addP.task.totalPrice" disabled />
                            </el-form-item>
                        </el-col>
                        <el-col :span="8">
                            <el-form-item label="已回款金额">
                                <el-input v-model="addP.task.paymentTotalPrice" disabled />
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="8">
                            <el-form-item label="金额" prop="money" v-if="addP.contract.payType == 1">
                                <el-input-number v-model="addP.model.money" :controls="false" :min="0"
                                    :max="addP.task.totalPrice - addP.task.paymentTotalPrice" style="width: 100%;" />
                            </el-form-item>
                            <el-form-item label="金额" prop="money" v-else>
                                <el-input-number v-model="addP.model.money" :controls="false" :min="0" :max="9999999999"
                                    style="width: 100%;" />
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-row>
                        <el-col :span="8">
                            <el-form-item label="回款日期" prop="paymentDate">
                                <el-date-picker v-model="addP.model.paymentDate" type="date" style="width: 100%;" />
                            </el-form-item>
                        </el-col>
                    </el-row>
                </div>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="addP.submit" :disabled="addP.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="editP.dialogVisible" title="回款记录编辑" width="75%" :show-close="false">
            <el-form :model="editP.model" label-width="100px" :rules="rules" ref="editPForm">
                <el-row v-if="addP.contract.isPreDeposit">
                    <el-col :span="8">
                        <el-form-item label="金额" prop="money">
                            <el-input-number v-model="editP.model.money" :controls="false" :min="0" :max="9999999999"
                                style="width: 100%;" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row v-else>
                    <el-col :span="8">
                        <el-form-item label="金额" prop="money" v-if="addP.contract.payType == 1">
                            <el-input-number v-model="editP.model.money" :controls="false" :min="0"
                                :max="editP.task.totalPrice - editP.task.paymentTotalPrice + editP.payment.money"
                                style="width: 100%;" />
                        </el-form-item>
                        <el-form-item label="金额" prop="money" v-else>
                            <el-input-number v-model="addP.model.money" :controls="false" :min="0" :max="9999999999"
                                style="width: 100%;" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="8">
                        <el-form-item label="回款日期" prop="paymentDate">
                            <el-date-picker v-model="editP.model.paymentDate" type="date" style="width: 100%;" />
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="editP.submit" :disabled="editP.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="final.dialogVisible" title="回款完成" width="50%" :show-close="false">
            <h1>是否确定该合同【{{ final.model.no }}】回款完成？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="final.submit" :disabled="final.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount, computed } from 'vue'
import { collectionStatusItems, invoiceTypeItems, isPreDepositItems, boolItems, payTypeItems } from '@/utils/magic'
import { queryContract, queryContracts } from '@/api/contract';
import { addInvoice, delInvoice, editInvoice } from '@/api/contract_invoice';
import { addPayment, editPayment, finalPayment } from '@/api/contract_payment';
import { message } from '@/components/divMessage/index'
import { reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addIForm = ref(null)
const editIForm = ref(null)
const addPForm = ref(null)
const editPForm = ref(null)
const rules = reactive({
    money: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    no: [
        { required: true, message: '请填写，不超过100个字！', trigger: 'blur' },
    ],
    taskID: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    paymentDate: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
})

const base = reactive({
    model: {
        no: "",
        employee: {
            name: "",
        },
        collectionStatus: 1,
        invoiceType: null,
        customer: {
            customerCompany: {
                name: "",
            },
            name: ""
        },
        startDate: new Date().getFullYear() + "-01-01",
        endDate: new Date().getFullYear() + "-12-31",
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
                prop: "office.name",
                label: "办事处",
                width: "5%",
            },
            {
                prop: "employee.name",
                label: "业务员",
                width: "5%",
            },
            {
                prop: "customer.customerCompany.name",
                label: "客户公司",
                width: "5%",
            },
            {
                prop: "customer.name",
                label: "客户",
                width: "5%",
            },
            {
                type: "transform",
                prop: "payType",
                items: payTypeItems,
                label: "付款类型",
                width: "5%",
            },
            {
                type: "boolean",
                prop: "isPreDeposit",
                label: "预存款",
                width: "5%",
            },
            {
                prop: "totalAmount",
                label: "总金额",
                width: "5%",
            },
            {
                prop: "paymentTotalAmount",
                label: "总回款",
                width: "5%",
            },
            {
                type: "contractNotPayment",
                prop: "notPaymentTotalAmount",
                label: "未回款",
                width: "5%",
            },
            {
                type: "transform",
                prop: "invoiceType",
                items: invoiceTypeItems,
                label: "发票",
                width: "5%",
            },
            {
                prop: "endPaymentDate",
                label: "完成时间",
                width: "5%",
            },
            {
                type: "transform",
                prop: "collectionStatus",
                items: collectionStatusItems,
                label: "状态",
                width: "5%",
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
                        label: "查看详情",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => {
                            if (row.isPreDeposit) {
                                return base.openViewPDialog(index, row)
                            } else {
                                base.openViewDialog(index, row)
                            }
                        }
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 2 && row.collectionStatus == 1) {
                                return true
                            }
                            return false
                        },
                        label: "添加发票",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openAddIDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 2 && row.collectionStatus == 1) {
                                return true
                            }
                            return false
                        },
                        label: "添加回款",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openAddPDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.status == 2 &&
                                row.collectionStatus == 1 &&
                                (row.isPreDeposit || row.totalAmount == row.paymentTotalAmount)) {
                                return true
                            }
                            return false
                        },
                        label: "回款完成",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.opemFinalDialog(index, row)
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
        queryContracts(base.model, base.pageData).then((res) => {
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
        queryContract(row.id).then((res) => {
            if (res.status == 1) {
                view.model = res.data
            }
        })

        view.dialogVisible = true
    },
    openViewPDialog: (index, row) => {
        queryContract(row.id).then((res) => {
            if (res.status == 1) {
                let arrs = units.paymentsSeparation(res.data.payments)
                res.data.payments = arrs[0]
                res.data.paymentAutos = arrs[1]
                viewP.model = res.data
            }
        })
        viewP.dialogVisible = true
    },
    openAddIDialog: (index, row) => {
        queryContract(row.id).then((res) => {
            if (res.status == 1) {
                addI.contract = res.data
            }
        })
        addI.model.contractID = row.id
        addI.dialogVisible = true
    },
    openAddPDialog: (index, row) => {
        addP.beforeFun(row.id)
        addP.task = {
            totalPrice: 0,
            paymentTotalPrice: 0,
        }
        addP.model = {
            taskID: null,
            money: 0,
            paymentDate: "",
        }
        addP.model.contractID = row.id
        addP.dialogVisible = true
    },
    opemFinalDialog: (index, row) => {
        final.model.contractID = row.id
        final.model.no = row.no
        final.dialogVisible = true
    }
})

const view = reactive({
    dialogVisible: false,
    model: {
        invoices: [],
        payments: [],
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
})

const viewP = reactive({
    dialogVisible: false,
    model: {
        invoices: [],
        payments: [],
        paymentAutos: [],
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
                prop: "money",
                label: "回款金额",
            },
        ]
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

})

const addI = reactive({
    dialogVisible: false,
    submitDisabled: false,
    payTypeString: computed(() => {
        var temp = "";
        payTypeItems.some((item) => {
            if (item.value == addI.contract.payType) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    invoiceTypeString: computed(() => {
        var temp = "";
        invoiceTypeItems.some((item) => {
            if (item.value == addI.contract.invoiceType) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    contract: {
        payType: null,
        invoiceType: null,
        totalAmount: 0,
        invoiceContent: "",
        invoices: [],
    },
    model: {
        contractID: null,
        no: "",
        money: 0,
    },
    column: {
        headers: [
            {
                prop: "createDate",
                label: "创建时间",
            },
            {
                prop: "no",
                label: "发票号",
            },
            {
                prop: "money",
                label: "金额",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => addI.openEditIDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "删除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => addI.openDelIDialog(index, row)
                    }
                ]
            },
        ]
    },
    submit: () => {
        addIForm.value.validate((valid) => {
            if (valid) {
                addI.submitDisabled = true
                addInvoice(addI.model).then((res) => {
                    if (res.status == 1) {
                        message("添加成功", "success")
                        queryContract(addI.model.contractID).then((res) => {
                            if (res.status == 1) {
                                addI.contract = res.data
                            }
                        })
                    } else {
                        message("添加失败", "error")
                    }
                    addI.model.no = ""
                    addI.model.money = 0
                    addI.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
    openEditIDialog: (index, row) => {
        editI.model.id = row.id
        editI.model.no = row.no
        editI.model.money = row.money
        editI.dialogVisible = true
    },
    openDelIDialog: (index, row) => {
        delI.model.id = row.id
        delI.model.no = row.no
        delI.dialogVisible = true
    },
})

const editI = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        no: "",
        money: 0,
    },
    submit: () => {
        editIForm.value.validate((valid) => {
            if (valid) {
                editI.submitDisabled = true
                editInvoice(editI.model).then((res) => {
                    if (res.status == 1) {
                        message("编辑成功", "success")
                        queryContract(addI.model.contractID).then((res) => {
                            if (res.status == 1) {
                                addI.contract = res.data
                            }
                        })
                    } else {
                        message("编辑失败", "error")
                    }
                    editI.dialogVisible = false
                    editI.model = {
                        id: null,
                        money: 0,
                        no: "",
                    }
                    editI.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

const delI = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 1,
        no: "",
    },
    submit: () => {
        delI.submitDisabled = true
        delInvoice(delI.model.id).then((res) => {
            if (res.status == 1) {
                message("删除成功", "success")
                queryContract(addI.model.contractID).then((res) => {
                    if (res.status == 1) {
                        addI.contract = res.data
                    }
                })
            } else {
                message("删除失败", "error")
            }
            delI.dialogVisible = false
            delI.model = {
                id: 1,
                no: "",
            }
            delI.submitDisabled = false
        })
    }
})

const addP = reactive({
    beforeFun: (contractID) => {
        queryContract(contractID).then((res) => {
            if (res.status == 1) {
                if (res.data.isPreDeposit) {
                    let arrs = units.paymentsSeparation(res.data.payments)
                    res.data.payments = arrs[0]
                    res.data.paymentAutos = arrs[1]
                }
                addP.contract = res.data
            }
        })
    },
    dialogVisible: false,
    submitDisabled: false,
    payTypeString: computed(() => {
        var temp = "";
        payTypeItems.some((item) => {
            if (item.value == addP.contract.payType) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    contract: {
        isPreDeposit: "",
        payType: null,
        totalAmount: 0,
        paymentContent: "",
        tasks: [],
        payments: [],
        paymentAutos: [],
    },
    task: {
        totalPrice: 0,
        paymentTotalPrice: 0,
    },
    taskChange: (val) => {
        addP.contract.tasks.some((item) => {
            if (item.id == val) {
                addP.task.totalPrice = item.totalPrice
                addP.task.paymentTotalPrice = item.paymentTotalPrice
                return;
            }
        });
        addP.model.money = 0
        addP.model.paymentDate = ""
    },
    model: {
        contractID: null,
        taskID: null,
        money: 0,
        paymentDate: "",
    },
    column: {
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
                type: "taskID",
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
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => addP.openEditPDialog(index, row)
                    }
                ]
            },
        ]
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
                prop: "money",
                label: "回款金额",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "编辑",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => addP.openEditPDialog(index, row)
                    }
                ]
            },
        ]
    },
    submit: () => {
        addPForm.value.validate((valid) => {
            if (valid) {
                addP.submitDisabled = true
                addPayment(addP.model).then((res) => {
                    if (res.status == 1) {
                        message("添加成功", "success")
                        addP.beforeFun(addP.model.contractID)
                        base.query()
                    } else {
                        message("添加失败", "error")
                    }
                    addP.task = {
                        totalPrice: 0,
                        paymentTotalPrice: 0,
                    }
                    addP.model.taskID = null
                    addP.model.money = 0
                    addP.model.paymentDate = 0
                    addP.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
    openEditPDialog: (index, row) => {
        editP.payment.money = row.money
        editP.task = row.task
        editP.model.id = row.id
        editP.model.contractID = row.contractID
        editP.model.taskID = row.taskID
        editP.model.money = 0
        editP.model.paymentDate = new Date(row.paymentDate.replace(/-/g, "/"))
        editP.dialogVisible = true
    },
})

const editP = reactive({
    dialogVisible: false,
    submitDisabled: false,
    payment: {
        money: 0,
    },
    task: {
        totalPrice: 0,
        paymentTotalPrice: 0,
    },
    model: {
        id: null,
        contractID: null,
        taskID: null,
        money: 0,
        paymentDate: "",
    },
    submit: () => {
        editPForm.value.validate((valid) => {
            if (valid) {
                editP.submitDisabled = true
                editPayment(editP.model).then((res) => {
                    if (res.status == 1) {
                        message("编辑成功", "success")
                        addP.beforeFun(addP.model.contractID)
                        base.query()
                    } else {
                        message("编辑失败", "error")
                    }
                    editP.payment = {
                        money: 0,
                    }
                    editP.task = {
                        totalPrice: 0,
                        paymentTotalPrice: 0,
                    }
                    editP.model = {
                        id: null,
                        contractID: null,
                        taskID: null,
                        money: 0,
                        paymentDate: "",
                    }
                    editP.dialogVisible = false
                    editP.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

const final = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        contractID: "",
        no: "",
    },
    submit: () => {
        final.submitDisabled = true
        finalPayment(final.model).then((res) => {
            if (res.status == 1) {
                message("操作成功", "success")
                base.query()
            } else {
                message("操作失败", "error")
            }
            final.model = {
                id: null,
                contractID: "",
                no: "",
            }
            final.dialogVisible = false
            final.submitDisabled = false
        })
    },
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
    base.query()
})
</script>