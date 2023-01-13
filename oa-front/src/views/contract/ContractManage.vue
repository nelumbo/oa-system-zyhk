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
                <el-select v-model="base.model.officeID" placeholder="办事处" clearable style="width: 100%;"
                    :disabled="!user().my.pids.includes('18')">
                    <el-option v-for="item in base.offices" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.employeeID" placeholder="业务员" clearable style="width: 100%;">
                    <el-option v-for="item in base.employees" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
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
                <el-select v-model="base.model.isSpecialNum" placeholder="特殊合同" clearable style="width: 100%;">
                    <el-option v-for="item in boolItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.isPreDepositNum" placeholder="预存款合同" clearable style="width: 100%;">
                    <el-option v-for="item in boolItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.payType" placeholder="付款类型" clearable style="width: 100%;">
                    <el-option v-for="item in payTypeItems" :key="item.value" :label="item.label" :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="5">
                <el-select v-model="base.model.invoiceType" placeholder="发票类型" clearable style="width: 100%;">
                    <el-option v-for="item in invoiceTypeItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
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
                <el-select v-model="base.model.havingInvoiceNum" placeholder="开票状态" clearable style="width: 100%;">
                    <el-option v-for="item in HavingInvoiceItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
        </el-row>
        <el-row :gutter="20">
            <el-col :span="5" :offset="1">
                <el-input v-model="base.model.employee.name" placeholder="业务员" clearable maxlength="25" />
            </el-col>
            <el-col :span="5">
                <el-input v-model="base.model.productName" placeholder="产品" clearable maxlength="25" />
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="view.dialogVisible" title="合同查看" width="90%" :show-close="false">
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

        <el-dialog v-model="approve.dialogVisible" title="合同审批" width="90%" :show-close="false">
            <el-divider content-position="left">
                <h2>基本信息</h2>
            </el-divider>
            <el-form :model="approve.model" label-width="120px">
                <el-row>
                    <el-col :span="6">
                        <el-form-item label="合同编号">
                            <el-input v-model="approve.model.no" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="录入时间">
                            <el-input v-model="approve.model.createDate" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-input v-model="approve.model.office.name" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="业务员">
                            <el-input v-model="approve.model.employee.name" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="区域">
                            <el-input v-model="approve.model.region.name" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="客户公司">
                            <el-input v-model="approve.model.customer.customerCompany.name" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="客户">
                            <el-input v-model="approve.model.customer.name" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6"></el-col>
                    <el-col :span="6">
                        <el-form-item label="签订日期">
                            <el-input v-model="approve.model.contractDate" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6" v-if="!approve.model.status || approve.model.status != 2">
                        <el-form-item label="交货日期">
                            <el-input v-model="approve.model.estimatedDeliveryDate" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="5" v-if="approve.model.status && approve.model.status == 2">
                        <el-form-item label="交货日期">
                            <el-input v-model="approve.model.estimatedDeliveryDate" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="1" v-if="approve.model.status && approve.model.status == 2">
                        <el-button type="primary" @click="approve.openEditEDDDidlog" :disabled="base.submitDisabled"
                            style="width:100%">修改</el-button>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="实际交货日期">
                            <el-input v-model="approve.model.endDeliveryDate" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="最后回款日期">
                            <el-input v-model="approve.model.endPaymentDate" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="特殊合同">
                            <el-input v-model="approve.isSpecialString" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="预存款合同">
                            <el-input v-model="approve.isPreDepositString" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="12"></el-col>
                    <el-col :span="6">
                        <el-form-item label="签订单位">
                            <el-input v-model="approve.model.vendor.name" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="付款类型">
                            <el-input v-model="approve.payTypeString" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="总金额">
                            <el-input v-model="approve.model.totalAmount" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="回款额">
                            <el-input v-model="approve.model.paymentTotalAmount" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="发票类型">
                            <el-input v-model="approve.invoiceTypeString" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="18"></el-col>
                    <el-col :span="24">
                        <el-form-item label="发票内容">
                            <el-input v-model="approve.model.invoiceContent" type="textarea"
                                :autosize="{ minRows: 1, maxRows: 9 }" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="24">
                        <el-form-item label="付款方式">
                            <el-input v-model="approve.model.paymentContent" type="textarea"
                                :autosize="{ minRows: 1, maxRows: 9 }" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="24">
                        <el-form-item label="备注">
                            <el-input v-model="approve.model.remark" type="textarea"
                                :autosize="{ minRows: 1, maxRows: 9 }" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="合同状态">
                            <el-input v-model="approve.statusString" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="生产状态">
                            <el-input v-model="approve.productionStatusString" readonly />
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="回款状态">
                            <el-input v-model="approve.collectionStatusString" readonly />
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <el-divider content-position="left" style="margin-top: 15px;">
                <h2>任务详情</h2>
            </el-divider>
            <divTable :columnObj="approve.column" :tableData="approve.model.tasks" :pageData="approve.pageData"
                :showAll="true" v-if="approve.model.tasks" />



            <el-row style="margin-top: 30px;" v-if="approve.model.status == 1">
                <el-col :span="2" :offset="7">
                    <el-button type="success" @click="approve.pass" :disabled="base.submitDisabled"
                        size="large">通过</el-button>
                </el-col>
                <el-col :span="2" :offset="6">
                    <el-button type="danger" @click="approve.reject" :disabled="base.submitDisabled"
                        size="large">驳回</el-button>
                </el-col>
            </el-row>

            <el-row style="margin-top: 30px;" v-if="approve.model.status == 2 || approve.model.status == 3">
                <el-col :span="2" :offset="4">
                    <el-button type="danger" @click="approve.openFinalDialog"
                        :disabled="approve.model.collectionStatus == 1 || (!approve.model.isPreDeposit && approve.model.productionStatus == 1)"
                        size="large">合同完成</el-button>
                </el-col>
                <el-col :span="2" :offset="4">
                    <el-button type="danger" @click="approve.openResetDialog"
                        :disabled="approve.model.collectionStatus != 2" size="large">回款状态回退</el-button>
                </el-col>
                <el-col :span="2" :offset="4">
                    <el-button type="danger" @click="approve.openRejectDialog"
                        :disabled="approve.model.status !== 2 && approve.model.status !== 3"
                        size="large">合同作废</el-button>
                </el-col>
            </el-row>

        </el-dialog>

        <el-dialog v-model="editEDD.dialogVisible" title="日期修改" width="50%" :show-close="false">
            <el-form :model="editEDD.model" label-width="100px" :rules="rules" ref="editEDDForm">
                <el-form-item label="交货日期" prop="estimatedDeliveryDate">
                    <el-date-picker v-model="editEDD.model.estimatedDeliveryDate" type="date" style="width: 100%;" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="editEDD.submit"
                            :disabled="editEDD.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="final.dialogVisible" title="预存款合同完成" width="50%" :show-close="false">
            <h1>是否确定该预存款合同已完成？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="final.submit" :disabled="final.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="reset.dialogVisible" title="合同回款状态回退" width="50%" :show-close="false">
            <h1>是否确定回退该合同的回款状态？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="reset.submit" :disabled="reset.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="reject.dialogVisible" title="合同作废" width="50%" :show-close="false">
            <h1>是否确定作废该合同？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="reject.submit" :disabled="reject.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="distribute.dialogVisible" title="分发" width="70%" :show-close="false">
            <el-form :model="distribute.model" label-width="140px" :rules="rules" ref="distributeForm">
                <el-row v-if="approve.model.isSpecial">
                    <el-form-item label="特殊提成百分比" prop="pushMoneyPercentages">
                        <el-input-number v-model="distribute.model.pushMoneyPercentages" :controls="false" :min="0" />
                    </el-form-item>
                </el-row>
                <el-row>
                    <el-col :span="24">
                        <el-form-item label="类型" prop="type">
                            <el-select v-model="distribute.model.type" @change="distribute.changeType"
                                style="width: 100%;">
                                <el-option v-for="item in taskTypeItems" :key="item.value" :label="item.label"
                                    :value="item.value" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row v-if="distribute.model.type == 3">
                    <el-col :span="2">
                        <el-form-item label="技术"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="distribute.office1" @change="distribute.changeOffice1"
                                style="width: 100%;">
                                <el-option v-for="item in distribute.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="technicianManID">
                            <el-select v-model="distribute.model.technicianManID" style="width: 100%;">
                                <el-option v-for="item in distribute.employees1" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="天数" prop="technicianDays">
                            <el-input-number v-model="distribute.model.technicianDays" :controls="false" :min="0" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row v-if="distribute.model.type > 1">
                    <el-col :span="2">
                        <el-form-item label="采购"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="distribute.office2" @change="distribute.changeOffice2"
                                style="width: 100%;">
                                <el-option v-for="item in distribute.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="purchaseManID">
                            <el-select v-model="distribute.model.purchaseManID" style="width: 100%;">
                                <el-option v-for="item in distribute.employees2" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="天数" prop="purchaseDays">
                            <el-input-number v-model="distribute.model.purchaseDays" :controls="false" :min="0" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="2">
                        <el-form-item label="仓库"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="distribute.office3" @change="distribute.changeOffice3"
                                style="width: 100%;">
                                <el-option v-for="item in distribute.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="inventoryManID">
                            <el-select v-model="distribute.model.inventoryManID" style="width: 100%;">
                                <el-option v-for="item in distribute.employees3" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="2">
                        <el-form-item label="物流"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="distribute.office4" @change="distribute.changeOffice4"
                                style="width: 100%;">
                                <el-option v-for="item in distribute.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="shipmentManID">
                            <el-select v-model="distribute.model.shipmentManID" style="width: 100%;">
                                <el-option v-for="item in distribute.employees4" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="distribute.submit"
                            :disabled="distribute.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="resetContractTask.dialogVisible" title="重置" width="70%" :show-close="false">
            <el-form :model="resetContractTask.model" label-width="140px" :rules="rules" ref="resetContractTaskForm">
                <el-row v-if="approve.model.isSpecial">
                    <el-form-item label="特殊提成百分比" prop="pushMoneyPercentages">
                        <el-input-number v-model="resetContractTask.model.pushMoneyPercentages" :controls="false"
                            :min="0" />
                    </el-form-item>
                </el-row>
                <el-row>
                    <el-col :span="24">
                        <el-form-item label="类型" prop="type">
                            <el-select v-model="resetContractTask.model.type" @change="resetContractTask.changeType"
                                style="width: 100%;">
                                <el-option v-for="item in taskTypeItems" :key="item.value" :label="item.label"
                                    :value="item.value" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row v-if="resetContractTask.model.type == 3">
                    <el-col :span="2">
                        <el-form-item label="技术"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="resetContractTask.office1" @change="resetContractTask.changeOffice1"
                                style="width: 100%;">
                                <el-option v-for="item in resetContractTask.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="technicianManID">
                            <el-select v-model="resetContractTask.model.technicianManID" style="width: 100%;">
                                <el-option v-for="item in resetContractTask.employees1" :key="item.id"
                                    :label="item.name" :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="天数" prop="technicianDays">
                            <el-input-number v-model="resetContractTask.model.technicianDays" :controls="false"
                                :min="0" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row v-if="resetContractTask.model.type > 1">
                    <el-col :span="2">
                        <el-form-item label="采购"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="resetContractTask.office2" @change="resetContractTask.changeOffice2"
                                style="width: 100%;">
                                <el-option v-for="item in resetContractTask.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="purchaseManID">
                            <el-select v-model="resetContractTask.model.purchaseManID" style="width: 100%;">
                                <el-option v-for="item in resetContractTask.employees2" :key="item.id"
                                    :label="item.name" :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="天数" prop="purchaseDays">
                            <el-input-number v-model="resetContractTask.model.purchaseDays" :controls="false"
                                :min="0" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="2">
                        <el-form-item label="仓库"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="resetContractTask.office3" @change="resetContractTask.changeOffice3"
                                style="width: 100%;">
                                <el-option v-for="item in resetContractTask.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="inventoryManID">
                            <el-select v-model="resetContractTask.model.inventoryManID" style="width: 100%;">
                                <el-option v-for="item in resetContractTask.employees3" :key="item.id"
                                    :label="item.name" :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="2">
                        <el-form-item label="物流"></el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="办事处">
                            <el-select v-model="resetContractTask.office4" @change="resetContractTask.changeOffice4"
                                style="width: 100%;">
                                <el-option v-for="item in resetContractTask.offices" :key="item.id" :label="item.name"
                                    :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                    <el-col :span="6">
                        <el-form-item label="员工" prop="shipmentManID">
                            <el-select v-model="resetContractTask.model.shipmentManID" style="width: 100%;">
                                <el-option v-for="item in resetContractTask.employees4" :key="item.id"
                                    :label="item.name" :value="item.id" />
                            </el-select>
                        </el-form-item>
                    </el-col>
                </el-row>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="resetContractTask.submit"
                            :disabled="resetContractTask.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>

        <el-dialog v-model="rejectContractTask.dialogVisible" title="驳回" width="50%" :show-close="false">
            <h1>是否确定驳回该任务？</h1>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="danger" @click="rejectContractTask.submit"
                            :disabled="rejectContractTask.submitDisabled">确定</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { ref, reactive, onBeforeMount, computed } from 'vue'
import {
    contractStatusItems, productionStatusItems, collectionStatusItems, payTypeItems,
    invoiceTypeItems, taskTypeItems, taskStatusItems,
    boolItems, HavingInvoiceItems
} from '@/utils/magic'
import { queryAllRegion } from "@/api/region"
import { editContractEDD, approveContract, finalContract, resetContract, rejectContract, queryContract, queryContracts } from "@/api/contract"
import { queryAllOffice } from "@/api/office"
import { queryAllEmployee } from "@/api/employee"
import { distributeTask, resetTask, rejectTask } from "@/api/contract_task"
import { message } from '@/components/divMessage/index'
import { reg_number, reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const distributeForm = ref(null)
const editEDDForm = ref(null)
const resetContractTaskForm = ref(null)
const rules = reactive({
    type: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    technicianManID: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    technicianDays: [
        { required: true, pattern: reg_number, message: '请输入有效数字', trigger: 'blur' }
    ],
    purchaseManID: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    purchaseDays: [
        { required: true, pattern: reg_number, message: '请输入有效数字', trigger: 'blur' }
    ],
    inventoryManID: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    shipmentManID: [
        { required: true, message: '请选择', trigger: 'blur' },
    ],
    estimatedDeliveryDate: [
        { required: true, message: '请选择交货日期', trigger: 'blur' },
    ],
    pushMoneyPercentages: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
})

const base = reactive({
    regions: [],
    offices: [],
    employee: [],
    model: {
        no: "",
        regionID: null,
        customer: {
            name: "",
            customerCompany: {
                name: "",
            }
        },
        employee: {
            name: "",
        },
        status: null,
        productionStatus: null,
        collectionStatus: null,
        payType: null,
        startDate: new Date().getFullYear() + "-01-01",
        endDate: new Date().getFullYear() + "-12-31",
        isSpecialNum: null,
        isPreDepositNum: null,
        havingInvoiceNum: null,
        productName: "",
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
                            if (user().my.pids.includes('19') && row.status > 0) {
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
    openAddDialog: () => {
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
    openApproveDialog: (index, row) => {
        queryContract(row.id).then((res) => {
            if (res.status == 1) {
                approve.model = res.data
            }
        })
        approve.dialogVisible = true
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
                width: "3%",
            },
            {
                prop: "product.type.name",
                label: "产品类型",
                width: "6%",
            },
            {
                prop: "product.name",
                label: "产品名称",
                width: "6%",
            },
            {
                prop: "product.specification",
                label: "规格",
                width: "5%",
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
                label: "标价(人民币)",
                width: "8%",
            },
            {
                prop: "productAttribute.standardPriceUSD",
                label: "标价(美元)",
                width: "8%",
            },
            {
                prop: "price",
                label: "售卖单价",
                width: "6%",
            },
            {
                prop: "totalPrice",
                label: "售卖总价",
                width: "6%",
            },
            {
                type: "employees",
                prop: "employees",
                label: "负责人",
                width: "9%",
            },
            {
                type: "taskStartDate",
                prop: "taskStartDate",
                label: "开始时间",
                width: "9%",
            },
            {
                type: "taskDays",
                prop: "taskDays",
                label: "限时天数",
                width: "6%",
            },
            {
                type: "taskFinalDate",
                prop: "taskFinalDate",
                label: "提交时间",
                width: "9%",
            },
            {
                type: "transform",
                prop: "status",
                label: "状态",
                items: taskStatusItems,
                width: "5%",
            },
            {
                type: "operation",
                label: "操作",
                width: "9%",
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

const approve = reactive({
    dialogVisible: false,
    submitDisabled: false,
    isSpecialString: computed(() => {
        if (approve.model.isSpecial) {
            return "是"
        } else {
            return "否"
        }
    }),
    isPreDepositString: computed(() => {
        if (approve.model.isPreDeposit) {
            return "是"
        } else {
            return "否"
        }
    }),
    payTypeString: computed(() => {
        var temp = "";
        payTypeItems.some((item) => {
            if (item.value == approve.model.payType) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    invoiceTypeString: computed(() => {
        var temp = "";
        invoiceTypeItems.some((item) => {
            if (item.value == approve.model.invoiceType) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    statusString: computed(() => {
        var temp = "";
        contractStatusItems.some((item) => {
            if (item.value == approve.model.status) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    productionStatusString: computed(() => {
        var temp = "";
        productionStatusItems.some((item) => {
            if (item.value == approve.model.productionStatus) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    collectionStatusString: computed(() => {
        var temp = "";
        collectionStatusItems.some((item) => {
            if (item.value == approve.model.collectionStatus) {
                temp = item.label;
                return;
            }
        });
        return temp;
    }),
    model: {
        id: null,
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
                prop: "product.numberCount",
                label: "库存数量",
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
                width: "8%",
            },
            {
                prop: "productAttribute.standardPriceUSD",
                label: "标准定价(美元)",
                width: "8%",
            },
            {
                prop: "price",
                label: "售卖单价",
                width: "8%",
            },
            {
                prop: "totalPrice",
                label: "售卖总价",
                width: "8%",
            },
            {
                prop: "remark",
                label: "业务员备注",
                width: "8%",
            },
            {
                prop: "employees",
                label: "负责人",
                width: "5%",
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
                            if (approve.model.status == 2 && row.status == 0) {
                                return true
                            }
                            return false
                        },
                        label: "分发",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => approve.openDistributeDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (approve.model.status == 2 && row.status > 0) {
                                return true
                            }
                            return false
                        },
                        label: "重置",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => approve.openResetTaskDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (approve.model.status == 2 && approve.model.isPreDeposit && row.status > -1) {
                                return true
                            }
                            return false
                        },
                        label: "驳回",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => approve.openRejectTaskDialog(index, row)
                    }
                ]
            },
        ]
    },
    submit: () => {
        approve.submitDisabled = true
        approveContract({ "id": approve.model.id, "isPass": approve.model.isPass }).then((res) => {
            if (res.status == 1) {
                message("审批成功", "success")
                queryContract(approve.model.id).then((res) => {
                    if (res.status == 1) {
                        approve.model = res.data
                    }
                })
            } else {
                message("审批失败", "error")
            }
            approve.submitDisabled = false
        })
    },
    pass: () => {
        approve.submitDisabled = true
        approveContract({ "id": approve.model.id, "isPass": true }).then((res) => {
            if (res.status == 1) {
                message("审批成功", "success")
                queryContract(approve.model.id).then((res) => {
                    if (res.status == 1) {
                        approve.model = res.data
                    }
                })
            } else {
                message("审批失败", "error")
            }
            approve.submitDisabled = false
        })
    },
    reject: () => {
        approve.submitDisabled = true
        approveContract({ "id": approve.model.id, "isPass": false }).then((res) => {
            if (res.status == 1) {
                message("审批成功", "success")
                queryContract(approve.model.id).then((res) => {
                    if (res.status == 1) {
                        approve.model = res.data
                    }
                })
            } else {
                message("审批失败", "error")
            }
            approve.model.isPass = null
            approve.submitDisabled = false
        })
    },
    openEditEDDDidlog: () => {
        editEDD.model.id = approve.model.id
        editEDD.model.estimatedDeliveryDate = ""
        editEDD.dialogVisible = true
    },
    openFinalDialog: () => {
        final.dialogVisible = true
    },
    openResetDialog: () => {
        reset.dialogVisible = true
    },
    openRejectDialog: () => {
        reject.dialogVisible = true
    },
    openDistributeDialog: (index, row) => {
        queryAllOffice().then((res) => {
            if (res.status == 1) {
                distribute.offices = res.data
            }
        })
        distribute.model.id = row.id
        distribute.model.contractID = row.contractID
        distribute.dialogVisible = true
    },
    openResetTaskDialog: (index, row) => {
        queryAllOffice().then((res) => {
            if (res.status == 1) {
                resetContractTask.offices = res.data
            }
        })
        resetContractTask.model.id = row.id
        resetContractTask.model.contractID = row.contractID
        resetContractTask.model.type = row.type
        resetContractTask.dialogVisible = true
    },
    openRejectTaskDialog: (index, row) => {
        rejectContractTask.model.id = row.id
        rejectContractTask.dialogVisible = true
    }
})

const editEDD = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        estimatedDeliveryDate: "",
    },
    submit: () => {
        editEDDForm.value.validate((valid) => {
            if (valid) {
                editEDD.submitDisabled = true
                editContractEDD(editEDD.model).then((res) => {
                    if (res.status == 1) {
                        message("编辑成功", "success")
                        approve.dialogVisible = false
                        base.query()
                    } else {
                        message("编辑失败", "error")
                    }
                    editEDD.dialogVisible = false
                    editEDD.model = {
                        id: null,
                        estimatedDeliveryDate: "",
                    }
                    editEDD.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const final = reactive({
    dialogVisible: false,
    submitDisabled: false,
    submit: () => {
        final.submitDisabled = true
        finalContract({ "id": approve.model.id }).then((res) => {
            if (res.status == 1) {
                message("提交成功", "success")
                approve.dialogVisible = false
                base.query()
            } else {
                message("提交失败", "error")
            }
            final.dialogVisible = false
            final.submitDisabled = false
        })
    }
})

const reset = reactive({
    dialogVisible: false,
    submitDisabled: false,
    submit: () => {
        reset.submitDisabled = true
        resetContract({ "id": approve.model.id }).then((res) => {
            if (res.status == 1) {
                message("回退成功", "success")
                approve.dialogVisible = false
                base.query()
            } else {
                message("回退失败", "error")
            }
            reset.dialogVisible = false
            reset.submitDisabled = false
        })
    }
})

const reject = reactive({
    dialogVisible: false,
    submitDisabled: false,
    submit: () => {
        reject.submitDisabled = true
        rejectContract({ "id": approve.model.id }).then((res) => {
            if (res.status == 1) {
                message("合同作废成功", "success")
                approve.dialogVisible = false
                base.query()
            } else {
                message("合同作废失败", "error")
            }
            reject.dialogVisible = false
            reject.submitDisabled = false
        })
    }
})

const distribute = reactive({
    dialogVisible: false,
    submitDisabled: false,
    offices: [],
    office1: null,
    office2: null,
    office3: null,
    office4: null,
    employees1: [],
    employees2: [],
    employees3: [],
    employees4: [],
    model: {
        id: null,
        contractID: null,
        pushMoneyPercentages: 0,
        type: 1,
        technicianManID: null,
        purchaseManID: null,
        inventoryManID: null,
        shipmentManID: null,
        technicianDays: 0,
        purchaseDays: 0,
    },
    submit: () => {
        distributeForm.value.validate((valid) => {
            if (valid) {
                distribute.submitDisabled = true
                distributeTask(distribute.model).then((res) => {
                    if (res.status == 1) {
                        message("分发成功", "success")
                        queryContract(approve.model.id).then((res) => {
                            if (res.status == 1) {
                                approve.model = res.data
                            }
                        })
                    } else {
                        message("分发失败", "error")
                    }
                    distribute.dialogVisible = false
                    distribute.model = {
                        id: null,
                        contractID: null,
                        technicianManID: null,
                        purchaseManID: null,
                        inventoryManID: null,
                        shipmentManID: null,
                        technicianDays: 0,
                        purchaseDays: 0,
                    }
                    distribute.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
    changeType: () => {
        distribute.office1 = null
        distribute.office2 = null
        distribute.office3 = null
        distribute.office4 = null
        distribute.employees1 = []
        distribute.employees2 = []
        distribute.employees3 = []
        distribute.employees4 = []
        distribute.model.technicianManID = null
        distribute.model.purchaseManID = null
        distribute.model.inventoryManID = null
        distribute.model.shipmentManID = null
    },
    changeOffice1: () => {
        distribute.employees1 = []
        distribute.model.technicianManID = null
        queryAllEmployee({ "office_id": distribute.office1 }).then((res) => {
            if (res.status == 1) {
                distribute.employees1 = res.data
            }
        })
    },
    changeOffice2: () => {
        distribute.employees2 = []
        distribute.model.purchaseManID = null
        queryAllEmployee({ "office_id": distribute.office2 }).then((res) => {
            if (res.status == 1) {
                distribute.employees2 = res.data
            }
        })
    },
    changeOffice3: () => {
        distribute.employees3 = []
        distribute.model.inventoryManID = null
        queryAllEmployee({ "office_id": distribute.office3 }).then((res) => {
            if (res.status == 1) {
                distribute.employees3 = res.data
            }
        })
    },
    changeOffice4: () => {
        distribute.employees4 = []
        distribute.model.shipmentManID = null
        queryAllEmployee({ "office_id": distribute.office4 }).then((res) => {
            if (res.status == 1) {
                distribute.employees4 = res.data
            }
        })
    },
})

const resetContractTask = reactive({
    dialogVisible: false,
    submitDisabled: false,
    offices: [],
    office1: null,
    office2: null,
    office3: null,
    office4: null,
    employees1: [],
    employees2: [],
    employees3: [],
    employees4: [],
    model: {
        id: null,
        contractID: null,
        pushMoneyPercentages: 0,
        type: 1,
        technicianManID: null,
        purchaseManID: null,
        inventoryManID: null,
        shipmentManID: null,
        technicianDays: 0,
        purchaseDays: 0,
    },
    submit: () => {
        resetContractTaskForm.value.validate((valid) => {
            if (valid) {
                resetContractTask.submitDisabled = true
                resetTask(resetContractTask.model).then((res) => {
                    if (res.status == 1) {
                        message("重置成功", "success")
                        queryContract(approve.model.id).then((res) => {
                            if (res.status == 1) {
                                approve.model = res.data
                            }
                        })
                    } else {
                        message("重置失败", "error")
                    }
                    resetContractTask.dialogVisible = false
                    resetContractTask.model = {
                        id: null,
                        contractID: null,
                        technicianManID: null,
                        purchaseManID: null,
                        inventoryManID: null,
                        shipmentManID: null,
                        technicianDays: 0,
                        purchaseDays: 0,
                    }
                    resetContractTask.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
    changeType: () => {
        resetContractTask.office1 = null
        resetContractTask.office2 = null
        resetContractTask.office3 = null
        resetContractTask.office4 = null
        resetContractTask.employees1 = []
        resetContractTask.employees2 = []
        resetContractTask.employees3 = []
        resetContractTask.employees4 = []
        resetContractTask.model.technicianManID = null
        resetContractTask.model.purchaseManID = null
        resetContractTask.model.inventoryManID = null
        resetContractTask.model.shipmentManID = null
    },
    changeOffice1: () => {
        resetContractTask.employees1 = []
        resetContractTask.model.technicianManID = null
        queryAllEmployee({ "office_id": resetContractTask.office1 }).then((res) => {
            if (res.status == 1) {
                resetContractTask.employees1 = res.data
            }
        })
    },
    changeOffice2: () => {
        resetContractTask.employees2 = []
        resetContractTask.model.purchaseManID = null
        queryAllEmployee({ "office_id": resetContractTask.office2 }).then((res) => {
            if (res.status == 1) {
                resetContractTask.employees2 = res.data
            }
        })
    },
    changeOffice3: () => {
        resetContractTask.employees3 = []
        resetContractTask.model.inventoryManID = null
        queryAllEmployee({ "office_id": resetContractTask.office3 }).then((res) => {
            if (res.status == 1) {
                resetContractTask.employees3 = res.data
            }
        })
    },
    changeOffice4: () => {
        resetContractTask.employees4 = []
        resetContractTask.model.shipmentManID = null
        queryAllEmployee({ "office_id": resetContractTask.office4 }).then((res) => {
            if (res.status == 1) {
                resetContractTask.employees4 = res.data
            }
        })
    },
})

const rejectContractTask = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
    },
    submit: () => {
        rejectContractTask.submitDisabled = true
        rejectTask(rejectContractTask.model).then((res) => {
            if (res.status == 1) {
                message("驳回成功", "success")
                queryContract(approve.model.id).then((res) => {
                    if (res.status == 1) {
                        approve.model = res.data
                    }
                })
            } else {
                message("驳回失败", "error")
            }
            rejectContractTask.dialogVisible = false
            rejectContractTask.model = {
                id: 0,
            }
            rejectContractTask.submitDisabled = false
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
    queryAllOffice().then((res) => {
        if (res.status == 1) {
            base.offices = res.data
        }
    })
    queryAllRegion().then((res) => {
        if (res.status == 1) {
            base.regions = res.data
        }
    })
    if (!user().my.pids.includes('18')) {
        base.model.officeID = Number(localStorage.getItem("officeID"))
    }
    base.query()
})
</script>