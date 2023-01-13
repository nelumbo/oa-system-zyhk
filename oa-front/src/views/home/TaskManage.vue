<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-input v-model="base.model.contract.no" placeholder="合同编号" clearable maxlength="50" />
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.status" placeholder="任务状态" clearable style="width: 100%;">
                    <el-option v-for="item in taskStatusSelectItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.productName" placeholder="产品名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.customerName" placeholder="客户名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-date-picker v-model="base.model.endDate" type="date" placeholder="结束时间" style="width: 100%;" />
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="view.dialogVisible" title="查看" width="75%" :show-close="false">
            <el-divider content-position="left">
                <h2>基本信息</h2>
            </el-divider>
            <el-form :model="view.model" label-width="120px">
                <el-form-item label="客户公司">
                    <el-input v-model="view.model.customer.customerCompany.name" readonly />
                </el-form-item>
                <el-form-item label="客户">
                    <el-input v-model="view.model.customer.name" readonly />
                </el-form-item>
                <el-form-item label="联系电话">
                    <el-input v-model="view.model.customer.phone" readonly />
                </el-form-item>
                <el-form-item label="合同备注">
                    <el-input v-model="view.model.remark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="产品库备注">
                    <el-input v-model="view.model.task.product.remark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="销售备注">
                    <el-input v-model="view.model.task.remark" type="textarea" autosize readonly />
                </el-form-item>
            </el-form>
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>负责人备注</h2>
            </el-divider>
            <el-form :model="view.model" label-width="120px">
                <el-form-item label="技术备注">
                    <el-input v-model="view.model.task.technicianRemark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="采购备注">
                    <el-input v-model="view.model.task.purchaseRemark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="仓库备注">
                    <el-input v-model="view.model.task.inventoryRemark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="装配备注">
                    <el-input v-model="view.model.task.assemblyRemark" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="物流备注">
                    <el-input v-model="view.model.task.shipmentRemark" type="textarea" autosize readonly />
                </el-form-item>
            </el-form>
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>任务详情</h2>
            </el-divider>
            <divTable :columnObj="view.column" :tableData="view.model.tasks" :allShow="true" />
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>采购详情</h2>
            </el-divider>
            <divTable :columnObj="view.columnP" :tableData="view.purchasings" :allShow="true" />
        </el-dialog>

        <el-dialog v-model="next.dialogVisible" title="提交" width="75%" :show-close="false">
            <el-divider content-position="left" style="margin-top: 50px;">
                <h2>采购详情</h2>
            </el-divider>
            <divTable :columnObj="next.column" :tableData="next.purchasings" :allShow="true" />
            <el-divider content-position="left" style="margin-top: 30px;">
                <h2>备注填写</h2>
            </el-divider>
            <el-form :model="next.model" label-width="60px">
                <el-form-item label="备注" prop="createRemark">
                    <el-input v-model="next.model.remark" type="textarea" :autosize="{ minRows: 9, maxRows: 18 }"
                        maxlength="300" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="next.submit" :disabled="next.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

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
            <el-divider content-position="left" style="margin-top: 30px;">
                <h2>需求清单</h2>
            </el-divider>
            <divTable :columnObj="add.columnP" :tableData="add.purchasings" :allShow="true" />
            <el-divider content-position="left" style="margin-top: 30px;">
                <h2>备注填写</h2>
            </el-divider>
            <el-form-item>
                <el-input v-model="add.task.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                    maxlength="300" />
            </el-form-item>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="add.submit" :disabled="add.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="addDLC.dialogVisible" title="添加" width="50%" :show-close="false">
            <el-divider content-position="left">
                <h2>产品信息</h2>
            </el-divider>
            <el-form :model="addDLC.product" label-width="100px">
                <el-form-item label="名称">
                    <el-input v-model.trim="addDLC.product.name" disabled />
                </el-form-item>
                <el-form-item label="型号">
                    <el-input v-model.trim="addDLC.product.version" disabled />
                </el-form-item>
                <el-form-item label="规格">
                    <el-input v-model.trim="addDLC.product.specification" disabled />
                </el-form-item>
                <el-form-item label="供应商">
                    <el-row>
                        <el-col :span="24" v-for="supplier in addDLC.product.suppliers" :key="supplier.id">
                            {{ supplier.name }}
                        </el-col>
                    </el-row>
                </el-form-item>
                <el-form-item label="库存数量">
                    <el-input v-model.trim="addDLC.product.numberCount" disabled />
                </el-form-item>
            </el-form>
            <el-divider content-position="left" style="margin-top: 30px;">
                <h2>采购信息</h2>
            </el-divider>
            <el-form :model="addDLC.model" label-width="100px" :rules="rules" ref="addForm">
                <el-form-item label="需求数量" prop="number">
                    <el-input-number v-model="addDLC.model.number" :controls="false" :min="1" :max="99999" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="addDLC.submit"
                            :disabled="addDLC.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="del.dialogVisible" title="删除" width="50%" :show-close="false">
            <h1>是否确定删除该条采购？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="del.submit" :disabled="del.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { ref, reactive, onBeforeMount } from 'vue'
import { taskStatusItems, taskStatusSelectItems } from '@/utils/magic'
import { queryMyTasks, queryMySavePurchasings } from "@/api/my"
import { queryContract } from "@/api/contract"
import { nextTask } from "@/api/contract_task"
import { queryProducts } from "@/api/product"
import { savePurchasing, submitPurchasing, delPurchasing, queryAllPurchasing } from "@/api/purchasing"
import { message } from '@/components/divMessage/index'
import { reg_number } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const rules = reactive({
    number: [
        { required: true, pattern: reg_number, message: '请输入大于零的有效数字', trigger: 'blur' },
    ],
})

const base = reactive({
    model: {
        contract: {
            no: "",
        },
        status: null,
        productName: "",
        endDate: "",
        customerName: "",
    },
    column: {
        headers: [
            {
                prop: "contract.no",
                label: "合同编号",
                width: "15%",
            },
            {
                prop: "product.name",
                label: "产品",
                width: "5%",
            },
            {
                prop: "number",
                label: "数量",
                width: "5%",
            },
            {
                prop: "product.numberCount",
                label: "库存",
                width: "5%",
            },
            {
                prop: "product.unit",
                label: "单位",
                width: "5%",
            },
            {
                prop: "contract.estimatedDeliveryDate",
                label: "合同交货时间",
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
                type: "taskStartDate",
                prop: "taskStartDate",
                label: "开始时间",
                width: "10%",
            },
            {
                type: "taskDays",
                prop: "taskDays",
                label: "限时天数",
                width: "10%",
            },
            {
                type: "taskFinalDate",
                prop: "taskFinalDate",
                label: "提交时间",
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

                            if (row.contract.status == 2) {
                                if (row.type == 3) {
                                    if (row.stauts == 2 && row.purchaseManID == user().my.id) {
                                        return true
                                    }
                                    if (row.stauts == 3 && row.inventoryManID == user().my.id) {
                                        return true
                                    }
                                    if (row.stauts == 4 && row.technicianManID == user().my.id) {
                                        return true
                                    }
                                    if (row.stauts == 5 && row.shipmentManID == user().my.id) {
                                        return true
                                    }
                                } else if (row.type == 2 || row.type == 1) {
                                    if (row.stauts == 3 && row.inventoryManID == user().my.id) {
                                        return true
                                    }
                                    if (row.stauts == 5 && row.shipmentManID == user().my.id) {
                                        return true
                                    }
                                }
                            }
                            return false
                        },
                        label: "提交|",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openNextDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.contract.status == 2) {
                                if (row.type == 3) {
                                    if (row.status == 1 && row.technicianManID == user().my.id) {
                                        return true
                                    }
                                } else if (row.type == 2) {
                                    if (row.status == 2 && row.purchaseManID == user().my.id) {
                                        return true
                                    }
                                }
                            }
                            return false
                        },
                        label: "提交||",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openAddDialog(index, row)
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
        queryMyTasks(base.model, base.pageData).then((res) => {
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
        queryContract(row.contractID).then((res) => {
            if (res.status == 1) {
                view.model = res.data
                view.model.task = row
            }
        })
        queryAllPurchasing({ "contractID": row.contractID, "taskID": row.id }).then((res) => {
            if (res.status == 1) {
                view.purchasings = res.data
            }
        })
        view.dialogVisible = true
    },
    openNextDialog: (index, row) => {
        queryAllPurchasing({ "contractID": row.contractID, "taskID": row.id }).then((res) => {
            if (res.status == 1) {
                next.purchasings = res.data
            }
        })
        next.model.id = row.id
        next.dialogVisible = true
    },
    openAddDialog: (index, row) => {
        add.task = row
        add.query()
        add.queryPurchasings()
        add.dialogVisible = true
    }
})

const view = reactive({
    dialogVisible: false,
    purchasings: [],
    model: {
        customer: {
            customerCompany: {
                name: "",
            },
            name: "",
            phone: "",
        },
        remark: "",
        product: {
            specification: "",
            remark: "",
        },
        tasks: [],
        task: {
            product: {
                name: "",
                remark: ""
            },
            remark: "",
        },
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
                type: "transform",
                prop: "status",
                label: "状态",
                items: taskStatusItems,
                width: "10%",
            },
        ]
    },
    columnP: {
        headers: [
            {
                prop: "id",
                label: "ID",
                width: "4%"
            },
            {
                prop: "createDate",
                label: "发起时间",
                width: "8%"
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
                width: "8%"
            },
            {
                prop: "number",
                label: "需求采购数量",
                width: "5%"
            },
            {
                prop: "realNumber",
                label: "实际采购数量",
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
            }
        ]
    },
})

const next = reactive({
    dialogVisible: false,
    submitDisabled: false,
    purchasings: [],
    column: {
        headers: [
            {
                prop: "id",
                label: "ID",
                width: "4%"
            },
            {
                prop: "createDate",
                label: "发起时间",
                width: "8%"
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
                width: "8%"
            },
            {
                prop: "number",
                label: "需求采购数量",
                width: "5%"
            },
            {
                prop: "realNumber",
                label: "实际采购数量",
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
            }
        ]
    },
    model: {
        id: 0,
        remark: "",
    },

    submit: () => {
        next.submitDisabled = true
        nextTask(next.model).then((res) => {
            if (res.status == 1) {
                message("提交成功", "success")
                base.query()
            } else {
                message("提交失败", "error")
            }
            next.dialogVisible = false
            next.model = {
                id: 0,
                remark: "",
            }
            next.submitDisabled = false
        })
    }
})

const add = reactive({
    submitDisabled: false,
    dialogVisible: false,
    task: {
        id: null,
        contractID: null,
        remark: "",
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
                prop: "numberCount",
                label: "库存数量",
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
    columnP: {
        headers: [
            {
                prop: "product.name",
                label: "名称",
            },
            {
                prop: "product.version",
                label: "型号",
            },
            {
                prop: "product.specification",
                label: "规格",
            },
            {
                prop: "number",
                label: "数量",
            },
            {
                prop: "product.unit",
                label: "单位",
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
                        onClick: (index, row) => add.openDelDialog(index, row)
                    }
                ]
            },
        ],
    },
    purchasings: [],
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
    submit: () => {
        add.submitDisabled = true
        nextTask({ "id": add.task.id, "remark": add.task.remark }).then((res) => {
            if (res.status == 1) {
                message("提交成功", "success")
                base.query()
            } else {
                message("提交失败", "error")
            }
            add.dialogVisible = false
            add.task = {
                "id": null,
                "contractID": null,
            }
            add.model = {
                name: "",
                version: "",
                specification: "",
            }
            add.submitDisabled = false
        })
        // submitPurchasing({ "contractID": add.task.contractID, "taskID": add.task.id }).then((res) => {
        //     if (res.status == 1) {
        //         message("提交成功", "success")
        //     } else {
        //         message("提交失败", "error")
        //     }
        //     add.dialogVisible = false
        //     add.task = {
        //         "id": null,
        //         "contractID": null,
        //     }
        //     add.model = {
        //         name: "",
        //         version: "",
        //         specification: "",
        //     }
        //     add.submitDisabled = false
        // })
    },
    queryPurchasings: () => {
        queryMySavePurchasings({ "contractID": add.task.contractID, "taskID": add.task.id }).then((res) => {
            if (res.status == 1) {
                add.purchasings = res.data
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
    openAddDLCDialog: (index, row) => {
        addDLC.product = row
        addDLC.model.contractID = add.task.contractID
        addDLC.model.taskID = add.task.id
        addDLC.model.productID = row.id
        addDLC.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.dialogVisible = true
    }
})

const addDLC = reactive({
    dialogVisible: false,
    submitDisabled: false,
    product: {
        name: "",
        version: "",
        specification: "",
        suppliers: [],
        numberCount: 0,
        number: 0,
    },
    model: {
        contractID: null,
        taskID: null,
        productID: null,
        number: 1,
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                addDLC.submitDisabled = true
                savePurchasing(addDLC.model).then((res) => {
                    if (res.status == 1) {
                        message("添加成功", "success")
                        add.queryPurchasings()
                    } else {
                        message("添加失败", "error")
                    }
                    addDLC.dialogVisible = false
                    addDLC.product = {
                        name: "",
                        version: "",
                        specification: "",
                        suppliers: [],
                        numberCount: 0,
                        number: 0,
                    }
                    addDLC.model = {
                        contractID: null,
                        taskID: null,
                        productID: null,
                        number: 1,
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
        id: null,
    },
    submit: () => {
        del.submitDisabled = true
        delPurchasing(del.model.id).then((res) => {
            if (res.status == 1) {
                message("删除成功", "success")
                add.queryPurchasings()
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
})
</script>