<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-select v-model="base.model.type" placeholder="报销类型" clearable style="width: 100%;">
                    <el-option v-for="item in expenseTypeItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                    <el-option v-for="item in expenseStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.employee.officeID" placeholder="办事处" clearable style="width: 100%;">
                    <el-option v-for="item in base.offices" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.employee.name" placeholder="员工" clearable maxlength="25" />
            </el-col>

            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
        </el-row>
        <el-row :gutter="20" style="margin-top: 5px;">
            <el-col :span="5" :offset="1">
                <el-date-picker v-model="base.model.startDate" type="date" placeholder="开始时间" style="width: 100%;" />
            </el-col>
            <el-col :span="5">
                <el-date-picker v-model="base.model.endDate" type="date" placeholder="结束时间" style="width: 100%;" />
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="view.dialogVisible" title="查看" width="50%" :show-close="false">
            <el-form :model="view.model" label-width="140px">
                <el-form-item label="报销状态">
                    <el-input v-model.trim="view.statusString" readonly />
                </el-form-item>
                <el-form-item label="报销类型">
                    <el-input v-model.trim="view.typeString" readonly />
                </el-form-item>
                <el-form-item label="发起时间">
                    <el-input v-model.trim="view.model.createDate" readonly />
                </el-form-item>
                <el-form-item label="办事处">
                    <el-input v-model.trim="view.model.employee.office.name" readonly />
                </el-form-item>
                <el-form-item label="员工">
                    <el-input v-model.trim="view.model.employee.name" readonly />
                </el-form-item>
                <el-form-item label="金额">
                    <el-input v-model.trim="view.model.amount" readonly />
                </el-form-item>
                <el-form-item label="申请理由">
                    <el-input v-model.trim="view.model.text" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="审核时间">
                    <el-input v-model.trim="view.model.approveDate" readonly />
                </el-form-item>
            </el-form>
        </el-dialog>

        <el-dialog v-model="approve.dialogVisible" title="审核" width="50%" :show-close="false">
            <el-form :model="approve.model" label-width="140px">
                <el-form-item label="员工补助金额">
                    <el-input v-model.trim="approve.model.employee.money" readonly />
                </el-form-item>
                <el-form-item label="办事处提成金额">
                    <el-input v-model.trim="approve.model.employee.office.money" readonly />
                </el-form-item>
                <el-form-item label="办事处业务费金额">
                    <el-input v-model.trim="approve.model.employee.office.businessMoney" readonly />
                </el-form-item>
                <el-divider />
                <el-form-item label="报销状态">
                    <el-input v-model.trim="approve.statusString" readonly />
                </el-form-item>
                <el-form-item label="报销类型">
                    <el-input v-model.trim="approve.typeString" readonly />
                </el-form-item>
                <el-form-item label="发起时间">
                    <el-input v-model.trim="approve.model.createDate" readonly />
                </el-form-item>
                <el-form-item label="办事处">
                    <el-input v-model.trim="approve.model.employee.office.name" readonly />
                </el-form-item>
                <el-form-item label="员工">
                    <el-input v-model.trim="approve.model.employee.name" readonly />
                </el-form-item>
                <el-form-item label="金额">
                    <el-input v-model.trim="approve.model.amount" readonly />
                </el-form-item>
                <el-form-item label="申请理由">
                    <el-input v-model.trim="approve.model.text" type="textarea" autosize readonly />
                </el-form-item>
                <el-form-item label="审核时间">
                    <el-input v-model.trim="approve.model.approveDate" readonly />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="approve.pass" :disabled="approve.submitDisabled"
                            style="margin-right: 250px;">通过</el-button>
                        <el-button type="danger" @click="approve.reject"
                            :disabled="approve.submitDisabled">拒绝</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { computed, reactive, onBeforeMount } from 'vue'
import { expenseTypeItems, expenseStatusItems } from '@/utils/magic'
import { approveExpense, queryExpenses } from "@/api/expense"
import { queryAllOffice } from "@/api/office";
import { message } from '@/components/divMessage/index'

import divTable from '../../components/divTable/index.vue'

const base = reactive({
    offices: [],
    model: {
        type: null,
        status: null,
        employee: {
            officeID: "",
            name: ""
        },
        startDate: "",
        endDate: "",
    },
    column: {
        headers: [
            {
                type: "transform",
                prop: "type",
                items: expenseTypeItems,
                label: "类型",
                width: "8%",
            },
            {
                prop: "createDate",
                label: "发起时间",
                width: "8%",
            },
            {
                prop: "employee.office.name",
                label: "办事处",
                width: "8%",
            },
            {
                prop: "employee.name",
                label: "员工",
                width: "8%",
            },
            {
                prop: "employee.phone",
                label: "电话",
                width: "8%",
            },
            {
                prop: "amount",
                label: "报销金额",
                width: "8%",
            },
            {
                prop: "approver.name",
                label: "审核",
                width: "8%",
            },
            {
                prop: "finance.name",
                label: "财务",
                width: "8%",
            },
            {
                prop: "cashier.name",
                label: "出纳",
                width: "8%",
            },
            {
                prop: "approveDate",
                label: "审核时间",
                width: "8%",
            },
            {
                type: "transform",
                prop: "status",
                items: expenseStatusItems,
                label: "状态",
                width: "8%",
            },
            {
                type: "operation",
                label: "操作",
                width: "12%",
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
                            if (row.status == 1) {
                                return true
                            } else if (row.status == 2) {
                                return true
                            } else if (row.status == 3) {
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
        queryExpenses(base.model, base.pageData).then((res) => {
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
        approve.model = row
        approve.dialogVisible = true
    },
    typeToString: (type) => {
        var temp = "";
        expenseTypeItems.some((item) => {
            if (item.value == type) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }
})

const view = reactive({
    dialogVisible: false,
    model: {
        id: null,
        employeeID: null,
        type: null,
        text: "",
        amount: "",
        status: "",
        createDate: "",
        approveDate: "",
        employee: {
            name: "",
            money: 0,
            office: {
                name: "",
                money: 0,
                businessMoney: 0,
            },
        },
    },
    statusString: computed(() => {
        var temp = "";
        expenseStatusItems.some((item) => {
            if (item.value == view.model.status) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    typeString: computed(() => {
        var temp = "";
        expenseTypeItems.some((item) => {
            if (item.value == view.model.type) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
})

const approve = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        employeeID: null,
        type: null,
        text: "",
        amount: "",
        status: "",
        createDate: "",
        approveDate: "",
        employee: {
            name: "",
            money: 0,
            office: {
                name: "",
                money: 0,
                businessMoney: 0,
            },
        },
        isPass: null,
    },
    statusString: computed(() => {
        var temp = "";
        expenseStatusItems.some((item) => {
            if (item.value == approve.model.status) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    typeString: computed(() => {
        var temp = "";
        expenseTypeItems.some((item) => {
            if (item.value == approve.model.type) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    submit: () => {
        approve.submitDisabled = true
        approveExpense(
            {
                "id": approve.model.id,
                "status": approve.model.status,
                "type": approve.model.type,
                "amount": approve.model.amount,
                "isPass": approve.model.isPass,
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
                employeeID: null,
                type: null,
                text: "",
                amount: "",
                status: "",
                createDate: "",
                employee: {
                    name: "",
                    money: 0,
                    office: {
                        name: "",
                        money: 0,
                        businessMoney: 0,
                    },
                },
                isPass: null,
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

onBeforeMount(() => {
    queryAllOffice().then((res) => {
        if (res.status == 1) {
            base.offices = res.data
        }
    })
    base.query()
})
</script>