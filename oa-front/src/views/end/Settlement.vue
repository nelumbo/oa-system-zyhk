<template>
    <div v-if="props.ice == 0">
        <el-row>
            <el-col :span="4" :offset="10">
                <el-button type="primary" size="large" @click="base.openStartDialog">开始结算</el-button>
            </el-col>
        </el-row>
    </div>

    <div v-if="props.ice == 1">
        <divTable :columnObj="run.column" :tableData="run.offices" :allShow="true" />
        <el-row style="margin-top: 10px;">
            <el-col :span="4" :offset="10">
                <el-button type="primary" size="large" @click="base.openEndDialog">结束结算</el-button>
            </el-col>
        </el-row>
    </div>

    <el-dialog v-model="start.dialogVisible" title="结算" width="50%" :show-close="false">
        <h1>是否开始年度结算？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="start.submit" :disabled="start.submitDisabled">确定</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="runDLC.dialogVisible" title="提交" width="50%" :show-close="false">
        <el-form :model="runDLC.model" label-width="120px" :rules="rules" ref="runDLCForm">
            <el-form-item label="结算提成" prop="lastYearMoney">
                <el-input-number v-model="runDLC.model.lastYearMoney" :controls="false" :min="-999999999" />
            </el-form-item>
            <el-form-item label="明年目标金额" prop="nextTaskLoad">
                <el-input-number v-model="runDLC.model.nextTaskLoad" :controls="false" :min="0" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="runDLC.submit" :disabled="runDLC.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>

    <el-dialog v-model="end.dialogVisible" title="结算" width="50%" :show-close="false">
        <h1>是否结束年度结算？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="end.submit" :disabled="end.submitDisabled">确定</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { ref, reactive, onBeforeMount } from 'vue'
import { iceStart, iceSubmit, iceEnd } from "@/api/settlement"
import { queryAllOffice } from "@/api/office"
import { message } from '@/components/divMessage/index'
import { reg_money } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const props = defineProps({
    ice: {
        type: Number,
        default: 2
    },
})

const runDLCForm = ref(null)
const rules = reactive({
    lastYearMoney: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
    nextTaskLoad: [
        { required: true, pattern: reg_money, message: '请输入最多三位小数的有效数字', trigger: 'blur' }
    ],
})

const base = reactive({
    openStartDialog: () => {
        start.dialogVisible = true
    },
    openEndDialog: () => {
        end.dialogVisible = true
    }
})

const start = reactive({
    dialogVisible: false,
    submitDisabled: false,
    submit: () => {
        start.submitDisabled = true
        iceStart().then((res) => {
            if (res.status == 1) {
                message("操作成功", "success")
                window.location.href = '/end'
            } else {
                message("操作成功", "error")
            }
            start.dialogVisible = false
            start.submitDisabled = false
        })
    },
})

const run = reactive({
    offices: [],
    column: {
        headers: [
            {
                prop: "number",
                label: "编号",
                width: "7%"
            },
            {
                prop: "name",
                label: "名称",
                width: "10%"
            },
            {
                prop: "businessMoney",
                label: "业务费",
                width: "10%"
            },
            {
                prop: "money",
                label: "可提成金额",
                width: "10%"
            },
            {
                prop: "moneyCold",
                label: "年底提成金额",
                width: "10%"
            },
            {
                prop: "taskLoad",
                label: "今年目标金额",
                width: "10%"
            },
            {
                prop: "targetLoad",
                label: "今年完成金额",
                width: "10%"
            },
            {
                prop: "lastYearMoney",
                label: "结算提成",
                width: "10%"
            },
            {
                prop: "nextTaskLoad",
                label: "明年目标金额",
                width: "10%"
            },
            {
                type: "operation",
                label: "操作",
                width: "20%",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.isSubmit) {
                                return false
                            }
                            return true
                        },
                        label: "提交",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => run.openRunDLCDialog(index, row)
                    },
                    {
                        isShow: (index, row) => {
                            if (row.isSubmit) {
                                return true
                            }
                            return false

                        },
                        label: "覆盖",
                        type: "success",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => run.openRunDLCDialog(index, row)
                    },
                ]
            },
        ],
    },
    query: () => {
        queryAllOffice().then((res) => {
            if (res.status == 1) {
                run.offices = res.data
            }
        })
    },
    openRunDLCDialog: (index, row) => {
        runDLC.model.id = row.id
        runDLC.model.lastYearMoney = row.lastYearMoney
        runDLC.model.nextTaskLoad = row.nextTaskLoad
        runDLC.dialogVisible = true
    }
})

const runDLC = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        lastYearMoney: 0,
        nextTaskLoad: 0,
    },
    submit: () => {
        runDLC.submitDisabled = true
        iceSubmit(runDLC.model).then((res) => {
            if (res.status == 1) {
                message("提交成功", "success")
                run.query()
            } else {
                message("提交失败", "error")
            }
            runDLC.dialogVisible = false
            runDLC.submitDisabled = false
        })
    },
})

const end = reactive({
    dialogVisible: false,
    submitDisabled: false,
    submit: () => {
        end.submitDisabled = true
        iceEnd().then((res) => {
            if (res.status == 1) {
                message("请求成功", "success")
                window.location.href = '/end'
            } else {
                message("请求失败", "error")
            }
            end.dialogVisible = false
            end.submitDisabled = false
        })
    },
})

onBeforeMount(() => {
    if (props.ice == 1) {
        run.query()
    }
})
</script>