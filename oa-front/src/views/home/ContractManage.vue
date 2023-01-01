<template>
    <div>
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
                <el-input v-model="base.model.customer.customerCompany.name" placeholder="客户单位" clearable
                    maxlength="25" />
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
                <el-button type="success" @click="base.openAddDialog">录入</el-button>
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
            <!-- <el-divider content-position="left" style="margin-top: 15px;">
                <h2>产品详情</h2>
            </el-divider>
            <el-divider content-position="left" style="margin-top: 15px;">
                <h2>开票记录</h2>
            </el-divider>
            <el-divider content-position="left" style="margin-top: 15px;">
                <h2>回款详情</h2>
            </el-divider> -->
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
    </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { ref, reactive, onBeforeMount, computed } from 'vue'
import { contractStatusItems, productionStatusItems, collectionStatusItems, payTypeItems, invoiceTypeItems } from '@/utils/magic'
import { queryAllRegion } from "@/api/region"
import { queryMyContracts } from "@/api/my"
import { delContract, queryContract } from "@/api/contract"
import { message } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const router = useRouter()
const delForm = ref(null)

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
        startDate: new Date().getFullYear() + "-01-01",
        endDate: new Date().getFullYear() + "-12-31",
        isSpecialNum: null,
        isPreDepositNum: null,
    },
    column: {
        headers: [
            {
                prop: "no",
                label: "合同编号",
                width: "13%",
            },
            {
                prop: "customer.customerCompany.name",
                label: "客户单位",
                width: "8%",
            },
            {
                prop: "customer.name",
                label: "客户",
                width: "8%",
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
                label: "预存款合同",
                width: "8%",
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
                width: "8%",
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
                width: "8%",
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
    openAddDialog: () => {
        router.push("entry")
    },
    openViewDialog: (index, row) => {
        queryContract(row.id).then((res) => {
            if (res.status == 1) {
                view.model = res.data
            }
        })
        view.dialogVisible = true
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

onBeforeMount(() => {
    queryAllRegion().then((res) => {
        if (res.status == 1) {
            base.regions = res.data
        }
    })
    base.query()
})
</script>