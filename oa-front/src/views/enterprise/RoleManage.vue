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

        <el-dialog v-model="add.dialogVisible" title="职位添加" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="60px" :rules="rules" ref="addForm">
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="add.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="权限">
                    <el-checkbox-group v-model="add.model.permissions">
                        <el-row>
                            <el-col :span="6" v-for="permission in add.permissions" :key="permission.id"
                                style="margin-bottom: 3px;">
                                <el-checkbox :label="permission">
                                    {{ permission.name }}
                                </el-checkbox>
                            </el-col>
                        </el-row>
                    </el-checkbox-group>
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

        <el-dialog v-model="edit.dialogVisible" title="职位编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="100px" :rules="rules" ref="editForm">
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="edit.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="权限">
                    <el-checkbox-group v-model="edit.model.permissions">
                        <el-row>
                            <el-col :span="6" v-for="permission in edit.permissions" :key="permission.id"
                                style="margin-bottom: 3px;">
                                <el-checkbox :label="permission">
                                    {{ permission.name }}
                                </el-checkbox>
                            </el-col>
                        </el-row>
                    </el-checkbox-group>
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

        <el-dialog v-model="del.dialogVisible" title="职位删除" width="50%" :show-close="false">
            <h1>是否确定删除【{{ del.model.name }}】？</h1>
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
import { ref, reactive, onBeforeMount } from 'vue'
import { addRole, delRole, editRole, queryRole, queryRoles } from "@/api/role"
import { queryAllPermission } from "@/api/permission"
import { message, messageForCRUD } from '@/components/divMessage/index'

import divTable from '@/components/divTable/index.vue'

const addForm = ref(null)
const editForm = ref(null)
const rules = reactive({
    name: [
        { required: true, message: '请输入职位名称', trigger: 'blur' },
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
                    {
                        isShow: (index, row) => {
                            return true
                        },
                        label: "删除",
                        type: "danger",
                        align: "center",
                        sortable: false,
                        onClick: (index, row) => base.openDelDialog(index, row)
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
        queryRoles(base.model, base.pageData).then((res) => {
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
        queryAllPermission().then((res) => {
            if (res.status == 1) {
                add.permissions = res.data
            }
        })
        add.dialogVisible = true
    },
    openEditDialog: (index, row) => {
        queryAllPermission().then((res) => {
            if (res.status == 1) {
                edit.permissions = res.data
            }
        })
        queryRole(row.id).then((res) => {
            if (res.status == 1) {
                edit.model = res.data
            }
        })
        edit.dialogVisible = true
    },
    openDelDialog: (index, row) => {
        del.model.id = row.id
        del.model.name = row.name
        del.dialogVisible = true
    }
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    permissions: [],
    model: {
        name: "",
        permissions: [],
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addRole(add.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add.model.name, "添加成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(add.model.name, "添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        name: "",
                        permissions: [],
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        });
    },
})

const edit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    permissions: [],
    model: {
        id: 0,
        name: "",
        permissions: [],
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editRole(edit.model).then((res) => {
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
                        permissions: [],
                    }
                    edit.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const del = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
        name: "",
    },
    submit: () => {
        del.submitDisabled = true
        delRole(del.model.id).then((res) => {
            if (res.status == 1) {
                messageForCRUD(del.model.name, "删除成功", "success")
                base.query()
            } else {
                messageForCRUD(del.model.name, "删除失败", "error")
            }
            del.dialogVisible = false
            del.model = {
                id: 0,
                name: "",
            }
            del.submitDisabled = false
        })
    }
})

onBeforeMount(() => {
    base.query()
})
</script>