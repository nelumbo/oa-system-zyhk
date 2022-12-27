<template>
    <div>
        <el-row :gutter="10">
            <el-col :span="6" :offset="8">
                <el-input v-model="base.model.name" placeholder="名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="1">
                <el-button type="primary" @click="base.query">查询</el-button>
            </el-col>
            <el-col :span="1">
                <el-button type="success" @click="base.openAddDialog">添加</el-button>
            </el-col>
        </el-row>
        <divTable :columnObj="base.column" :tableData="base.tableData" :pageData="base.pageData"
            :handleSizeChange="base.handleSizeChange" :handleCurrentChange="base.handleCurrentChange" />

        <el-dialog v-model="add.dialogVisible" title="区域添加" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="60px" :rules="rules" ref="addForm">
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="add.model.name" maxlength="50" />
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

        <el-dialog v-model="edit.dialogVisible" title="区域编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="100px" :rules="rules" ref="editForm">
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="edit.model.name" maxlength="50" />
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
    </div>
</template>

<script setup>
import { ref, reactive, onBeforeMount } from 'vue'
import { addRegion, editRegion, queryRegions } from "@/api/region";
import { message, messageForCRUD } from '@/components/divMessage/index'

import divTable from '../../components/divTable/index.vue'

const addForm = ref(null)
const editForm = ref(null)
const rules = reactive({
    name: [
        { required: true, message: '请输入区域名称', trigger: 'blur' },
    ],
})
const base = reactive({
    model: {
        name: "",
    },
    column: {
        headers: [
            {
                prop: "id",
                label: "ID",
            },
            {
                prop: "name",
                label: "名称",
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
        queryRegions(base.model, base.pageData).then((res) => {
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
        edit.dialogVisible = true
    }
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        name: "",
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addRegion(add.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add.model.name, "添加成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(add.model.name, "添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        name: "",
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
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editRegion(edit.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(edit.model.name, "编辑成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(edit.model.name, "编辑失败", "error")
                    }
                    edit.dialogVisible = false
                    edit.model = {
                        id: 0,
                        name: "",
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