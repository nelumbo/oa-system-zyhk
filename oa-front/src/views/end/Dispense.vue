<template>
    <el-row>
        <el-col :span="6" :offset="3">
            年度提成额度：<el-input v-model="props.employee.office.lastYearMoney" disabled />
        </el-col>
        <el-col :span="6" :offset="3">
            目前分配额度：<el-input v-model="base.yearMoneyCount" disabled />
        </el-col>
    </el-row>
    <el-divider content-position="left">
    </el-divider>
    <el-row v-for="employee in base.employees">
        <el-col :span="6" :offset="3">
            员工：<el-input v-model="employee.name" disabled />
        </el-col>
        <el-col :span="6" :offset="3">
            年度提成： <el-input-number v-model="employee.yearMoney" :controls="false" :min="0" style="width:100%" />
        </el-col>
    </el-row>
    <el-row style="margin-top: 30px;">
        <el-col :span="4" :offset="10">
            <el-button type="primary" size="large" @click="base.openBaseDLCDialog"
                :disabled="base.yearMoneyCount != props.employee.office.lastYearMoney"
                v-if="props.employee.office.isSetSubmit == 0 || props.employee.office.isSetSubmit == -1">提交</el-button>
            <el-button type="primary" size="large" @click="base.openBaseDLCDialog"
                :disabled="base.yearMoneyCount != props.employee.office.lastYearMoney"
                v-if="props.employee.office.isSetSubmit == 1">重新提交</el-button>
        </el-col>
    </el-row>

    <el-dialog v-model="baseDLC.dialogVisible" title="结算" width="50%" :show-close="false">
        <h1>是否提交年度提成分配方案？</h1>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="danger" @click="baseDLC.submit"
                        :disabled="base.yearMoneyCount != props.employee.office.lastYearMoney">确定</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { reactive, onBeforeMount, computed } from 'vue'
import { settSubmit } from "@/api/settlement"
import { queryAllEmployee } from "@/api/employee"
import { message } from '@/components/divMessage/index'

const props = defineProps({
    employee: {
        type: Object,
        default: {
            money: 0,
            officeID: 0,
            office: {
                lastYearMoney: 0,
                businessMoney: 0,
                money: 0,
                moneyCold: 0,
                isSetSubmit: 0,
            },
        }
    }
})

const base = reactive({
    employees: [],
    yearMoneyCount: computed(() => {
        if (base.employees) {
            let temp = 0
            base.employees.forEach((item, index) => {
                temp += item.yearMoney
            })
            return temp
        } else {
            return 0
        }
    }),
    query: () => {
        if (props.employee.officeID != 0) {
            queryAllEmployee({ "officeID": props.employee.officeID }).then((res) => {
                if (res.status == 1) {
                    base.employees = res.data
                }
            })
        }
    },
    openBaseDLCDialog: () => {
        baseDLC.dialogVisible = true
    },
})

const baseDLC = reactive({
    dialogVisible: false,
    submitDisabled: false,
    submit: () => {
        baseDLC.submitDisabled = true
        settSubmit(base.employees).then((res) => {
            if (res.status == 1) {
                message("请求成功", "success")
            } else {
                message("请求失败", "error")
            }
            window.location.href = '/end'
            baseDLC.dialogVisible = false
            baseDLC.submitDisabled = false
        })
    },
})

onBeforeMount(() => {
    base.query()
})
</script>