<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="6" :offset="8">
                <el-select v-model="base.model.status" placeholder="状态" clearable style="width: 100%;">
                    <el-option v-for="item in predesignTaskStatusItems" :key="item.value" :label="item.label"
                        :value="item.value" />
                </el-select>
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="submit.dialogVisible" title="提交" width="50%" :show-close="false">
            <el-form :model="submit.model" label-width="60px">
                <el-form-item label="备注">
                    <el-input v-model="submit.model.finalRemark" type="textarea" :autosize="{ minRows: 9, maxRows: 19 }"
                        maxlength="300" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <div style="text-align: center;">
                        <el-button type="primary" @click="submit.submit"
                            :disabled="submit.submitDisabled">提交</el-button>
                    </div>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { predesignTaskStatusItems } from '@/utils/magic'
import { queryMyPredesignTasks } from "@/api/my"
import { submitPredesignTask } from "@/api/predesign_task"
import { message } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    model: {
        status: null,
    },
    column: {
        headers: [
            {
                prop: "endDate",
                label: "预期完成时间",
            },
            {
                type: "textarea",
                prop: "createRemark",
                label: "设计要求",
            },
            {
                prop: "finalDate",
                label: "提交日期",
            },
            {
                type: "transform",
                prop: "status",
                items: predesignTaskStatusItems,
                label: "状态",
            },
            {
                type: "operation",
                label: "操作",
                operations: [
                    {
                        isShow: (index, row) => {
                            if (row.status == 1) {
                                return true
                            }
                            return false
                        },
                        label: "提交",
                        type: "primary",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openSubmitDialog(index, row)
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
        queryMyPredesignTasks(base.model, base.pageData).then((res) => {
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
    openSubmitDialog: (index, row) => {
        submit.model.id = row.id
        submit.dialogVisible = true
    },
})

const submit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: null,
        finalRemark: "",
    },
    submit: () => {
        submit.submitDisabled = true
        submitPredesignTask(submit.model).then((res) => {
            if (res.status == 1) {
                message("发起成功", "success")
                base.query()
            } else {
                message("发起失败", "error")
            }
            submit.dialogVisible = false
            submit.model = {
                id: null,
                finalRemark: "",
            }
            submit.submitDisabled = false
        })

    },
})

onBeforeMount(() => {
    base.query()
})
</script>