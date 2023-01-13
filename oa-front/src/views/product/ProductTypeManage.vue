<template>
    <el-row :gutter="20">
        <el-col :span="6" :offset="8">
            <el-input v-model="base.model.name" placeholder="名称" clearable maxlength="25" />
        </el-col>
        <el-col :span="1">
            <el-button type="primary" @click="base.query">查询</el-button>
        </el-col>
        <el-col :span="1" v-if="user().my.pids.includes('37')">
            <el-button type="success" @click="base.openAddDialog">添加</el-button>
        </el-col>
    </el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
        :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

    <el-dialog v-model="add.dialogVisible" title="添加" width="50%" :show-close="false">
        <el-form :model="add.model" label-width="160px" :rules="rules" ref="addForm">
            <el-form-item label="名称" prop="name">
                <el-input v-model.trim="add.model.name" maxlength="50" />
            </el-form-item>
            <el-form-item label="标准提成百分比" prop="pushMoneyPercentages">
                <el-input-number v-model="add.model.pushMoneyPercentages" :controls="false" :min="-100" :max="100" />
            </el-form-item>
            <el-form-item label="最低提成百分比" prop="minPushMoneyPercentages">
                <el-input-number v-model="add.model.minPushMoneyPercentages" :controls="false" :min="-100" :max="100" />
            </el-form-item>
            <el-form-item label="高价提成百分比" prop="pushMoneyPercentagesUp">
                <el-input-number v-model="add.model.pushMoneyPercentagesUp" :controls="false" :min="-100" :max="100" />
            </el-form-item>
            <el-form-item label="低价提成下降百分比" prop="pushMoneyPercentagesDown">
                <el-input-number v-model="add.model.pushMoneyPercentagesDown" :controls="false" :min="-100"
                    :max="100" />
            </el-form-item>
            <el-form-item label="标准业务费用百分比" prop="businessMoneyPercentages">
                <el-input-number v-model="add.model.businessMoneyPercentages" :controls="false" :min="-100"
                    :max="100" />
            </el-form-item>
            <el-form-item label="高价业务费用百分比" prop="businessMoneyPercentagesUp">
                <el-input-number v-model="add.model.businessMoneyPercentagesUp" :controls="false" :min="-100"
                    :max="100" />
            </el-form-item>
            <el-form-item label="是否计算任务量">
                <el-radio-group v-model="add.model.isTaskLoad">
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

    <el-dialog v-model="edit.dialogVisible" title="编辑" width="50%" :show-close="false">
        <el-form :model="edit.model" label-width="160px" :rules="rules" ref="editForm">
            <el-form-item label="名称" prop="name">
                <el-input v-model.trim="edit.model.name" maxlength="50" />
            </el-form-item>
            <el-form-item label="标准提成百分比" prop="pushMoneyPercentages">
                <el-input-number v-model="edit.model.pushMoneyPercentages" :controls="false" :min="-100" :max="100" />
            </el-form-item>
            <el-form-item label="最低提成百分比" prop="minPushMoneyPercentages">
                <el-input-number v-model="edit.model.minPushMoneyPercentages" :controls="false" :min="-100"
                    :max="100" />
            </el-form-item>
            <el-form-item label="高价提成百分比" prop="pushMoneyPercentagesUp">
                <el-input-number v-model="edit.model.pushMoneyPercentagesUp" :controls="false" :min="-100" :max="100" />
            </el-form-item>
            <el-form-item label="低价提成下降百分比" prop="pushMoneyPercentagesDown">
                <el-input-number v-model="edit.model.pushMoneyPercentagesDown" :controls="false" :min="-100"
                    :max="100" />
            </el-form-item>
            <el-form-item label="标准业务费用百分比" prop="businessMoneyPercentages">
                <el-input-number v-model="edit.model.businessMoneyPercentages" :controls="false" :min="-100"
                    :max="100" />
            </el-form-item>
            <el-form-item label="高价业务费用百分比" prop="businessMoneyPercentagesUp">
                <el-input-number v-model="edit.model.businessMoneyPercentagesUp" :controls="false" :min="-100"
                    :max="100" />
            </el-form-item>
            <el-form-item label="是否计算任务量">
                <el-radio-group v-model="edit.model.isTaskLoad">
                    <el-radio :label="true">是</el-radio>
                    <el-radio :label="false">否</el-radio>
                </el-radio-group>
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <div style="text-align: center;">
                    <el-button type="primary" @click="edit.submit" :disabled="edit.submitDisabled">提交</el-button>
                </div>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
import { user } from '@/pinia/modules/user'
import { ref, reactive, onBeforeMount } from 'vue'
import { addProductType, editProductType, queryProductTypes } from "@/api/product_type"
import { message } from '@/components/divMessage/index'
import { reg_number_2d } from '@/utils/regex'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const editForm = ref(null)
const rules = reactive({
    name: [
        { required: true, message: '请输入类型名称', trigger: 'blur' },
    ],
    pushMoneyPercentages: [
        { required: true, pattern: reg_number_2d, message: '请输入最多两位小数的有效数字', trigger: 'blur' }
    ],
    minPushMoneyPercentages: [
        { required: true, pattern: reg_number_2d, message: '请输入最多两位小数的有效数字', trigger: 'blur' }
    ],
    pushMoneyPercentagesUp: [
        { required: true, pattern: reg_number_2d, message: '请输入最多两位小数的有效数字', trigger: 'blur' }
    ],
    pushMoneyPercentagesDown: [
        { required: true, pattern: reg_number_2d, message: '请输入最多两位小数的有效数字', trigger: 'blur' }
    ],
    businessMoneyPercentages: [
        { required: true, pattern: reg_number_2d, message: '请输入最多两位小数的有效数字', trigger: 'blur' }
    ],
    businessMoneyPercentagesUp: [
        { required: true, pattern: reg_number_2d, message: '请输入最多两位小数的有效数字', trigger: 'blur' }
    ],
})

const base = reactive({
    model: {
        name: "",
    },
    column: {
        headers: [
            {
                prop: "name",
                label: "名称",
            },
            {
                prop: "pushMoneyPercentages",
                label: "标准提成百分比",
            },
            {
                prop: "minPushMoneyPercentages",
                label: "最低提成百分比",
            },
            {
                prop: "pushMoneyPercentagesUp",
                label: "高价提成百分比",
            },
            {
                prop: "pushMoneyPercentagesDown",
                label: "低价提成下降百分比",
            },
            {
                prop: "businessMoneyPercentages",
                label: "标准业务费用百分比",
            },
            {
                prop: "businessMoneyPercentagesUp",
                label: "高价业务费用百分比",
            },
            {
                type: "boolean",
                prop: "isTaskLoad",
                label: "是否计算任务量",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (user().my.pids.includes('38')) {
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
        queryProductTypes(base.model, base.pageData).then((res) => {
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
    openEditDialog: (index, row) => {
        edit.model.id = row.id
        edit.model.name = row.name
        edit.model.pushMoneyPercentages = row.pushMoneyPercentages
        edit.model.minPushMoneyPercentages = row.minPushMoneyPercentages
        edit.model.pushMoneyPercentagesUp = row.pushMoneyPercentagesUp
        edit.model.pushMoneyPercentagesDown = row.pushMoneyPercentagesDown
        edit.model.businessMoneyPercentages = row.businessMoneyPercentages
        edit.model.businessMoneyPercentagesUp = row.businessMoneyPercentagesUp
        edit.model.isTaskLoad = row.isTaskLoad
        edit.dialogVisible = true
    },
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        name: "",
        pushMoneyPercentages: 0,
        minPushMoneyPercentages: 0,
        pushMoneyPercentagesUp: 0,
        pushMoneyPercentagesDown: 0,
        businessMoneyPercentages: 0,
        businessMoneyPercentagesUp: 0,
        isTaskLoad: true,
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addProductType(add.model).then((res) => {
                    if (res.status == 1) {
                        message("添加成功", "success")
                        base.query()
                    } else {
                        message("添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        name: "",
                        pushMoneyPercentages: 0,
                        minPushMoneyPercentages: 0,
                        pushMoneyPercentagesUp: 0,
                        pushMoneyPercentagesDown: 0,
                        businessMoneyPercentages: 0,
                        businessMoneyPercentagesUp: 0,
                        isTaskLoad: true,
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const edit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
        name: "",
        pushMoneyPercentages: 0,
        minPushMoneyPercentages: 0,
        pushMoneyPercentagesUp: 0,
        pushMoneyPercentagesDown: 0,
        businessMoneyPercentages: 0,
        businessMoneyPercentagesUp: 0,
        isTaskLoad: true,
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editProductType(edit.model).then((res) => {
                    if (res.status == 1) {
                        message("编辑成功", "success")
                        base.query()
                    } else {
                        message("编辑失败", "error")
                    }
                    edit.dialogVisible = false
                    edit.model = {
                        id: 0,
                        name: "",
                        pushMoneyPercentages: 0,
                        minPushMoneyPercentages: 0,
                        pushMoneyPercentagesUp: 0,
                        pushMoneyPercentagesDown: 0,
                        businessMoneyPercentages: 0,
                        businessMoneyPercentagesUp: 0,
                        isTaskLoad: true,
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

onBeforeMount(() => {
    base.query()
})
</script>