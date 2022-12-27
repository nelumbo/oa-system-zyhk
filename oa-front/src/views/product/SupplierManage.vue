<template>
    <div>
        <el-row :gutter="10">
            <el-col :span="6" :offset="2">
                <el-input v-model="base.model.name" placeholder="供应商名称" clearable maxlength="25" />
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.linkman" placeholder="联系人" clearable maxlength="25" />
            </el-col>
            <el-col :span="6">
                <el-input v-model="base.model.phone" placeholder="联系电话" clearable maxlength="25" />
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

        <el-dialog v-model="add.dialogVisible" title="供应商添加" width="50%" :show-close="false">
            <el-form :model="add.model" label-width="100px" :rules="rules" ref="addForm">
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="add.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="地址">
                    <el-input v-model.trim="add.model.address" maxlength="100" />
                </el-form-item>
                <el-form-item label="网站">
                    <el-input v-model.trim="add.model.web" maxlength="50" />
                </el-form-item>
                <el-form-item label="联系人">
                    <el-input v-model.trim="add.model.linkman" maxlength="50" />
                </el-form-item>
                <el-form-item label="联系电话">
                    <el-input v-model.trim="add.model.phone" maxlength="50" />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model.trim="add.model.wechatID" maxlength="50" />
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model.trim="add.model.email" maxlength="50" />
                </el-form-item>
                <el-form-item label="主营产品概述">
                    <el-input v-model="add.model.description" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="add.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
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

        <el-dialog v-model="view.dialogVisible" title="供应商查看" width="50%" :show-close="false">
            <el-form :model="view.model" label-width="100px">
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="view.model.name" disabled />
                </el-form-item>
                <el-form-item label="地址">
                    <el-input v-model.trim="view.model.address" disabled />
                </el-form-item>
                <el-form-item label="网站">
                    <el-input v-model.trim="view.model.web" disabled />
                </el-form-item>
                <el-form-item label="联系人">
                    <el-input v-model.trim="view.model.linkman" disabled />
                </el-form-item>
                <el-form-item label="联系电话">
                    <el-input v-model.trim="view.model.phone" disabled />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model.trim="view.model.wechatID" disabled />
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model.trim="view.model.email" disabled />
                </el-form-item>
                <el-form-item label="主营产品概述">
                    <el-input v-model="view.model.description" type="textarea" :autosize=true disabled />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="view.model.remark" type="textarea" :autosize=true disabled />
                </el-form-item>
            </el-form>
        </el-dialog>

        <el-dialog v-model="edit.dialogVisible" title="供应商编辑" width="50%" :show-close="false">
            <el-form :model="edit.model" label-width="120px" :rules="rules" ref="editForm">
                <el-form-item label="名称" prop="name">
                    <el-input v-model.trim="edit.model.name" maxlength="50" />
                </el-form-item>
                <el-form-item label="地址">
                    <el-input v-model.trim="edit.model.address" maxlength="100" />
                </el-form-item>
                <el-form-item label="网站">
                    <el-input v-model.trim="edit.model.web" maxlength="50" />
                </el-form-item>
                <el-form-item label="联系人">
                    <el-input v-model.trim="edit.model.linkman" maxlength="50" />
                </el-form-item>
                <el-form-item label="联系电话">
                    <el-input v-model.trim="edit.model.phone" maxlength="50" />
                </el-form-item>
                <el-form-item label="微信号">
                    <el-input v-model.trim="edit.model.wechatID" maxlength="50" />
                </el-form-item>
                <el-form-item label="邮箱">
                    <el-input v-model.trim="edit.model.email" maxlength="50" />
                </el-form-item>
                <el-form-item label="主营产品概述">
                    <el-input v-model="edit.model.description" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
                </el-form-item>
                <el-form-item label="备注">
                    <el-input v-model="edit.model.remark" type="textarea" :autosize="{ minRows: 3, maxRows: 9 }"
                        maxlength="300" />
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
import { addSupplier, editSupplier, querySuppliers } from "@/api/supplier"
import { message, messageForCRUD } from '@/components/divMessage/index'

import divTable from '../../components/divTable/index.vue'

const addForm = ref(null)
const editForm = ref(null)
const rules = reactive({
    name: [
        { required: true, message: '请输入供应商名称', trigger: 'blur' },
    ],
})

const base = reactive({
    model: {
        name: "",
        linkman: "",
        phone: "",
    },
    column: {
        headers: [
            {
                prop: "name",
                label: "名称",
            },
            {
                prop: "web",
                label: "网站",
            },
            {
                prop: "linkman",
                label: "联系人",
            },
            {
                prop: "phone",
                label: "联系电话",
            },
            {
                type: "operation",
                label: "操作",
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
        querySuppliers(base.model, base.pageData).then((res) => {
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
    openViewDialog: (index, row) => {
        view.model.name = row.name
        view.model.address = row.address
        view.model.web = row.web
        view.model.linkman = row.linkman
        view.model.phone = row.phone
        view.model.wechatID = row.wechatID
        view.model.email = row.email
        view.model.description = row.description
        view.model.remark = row.remark
        view.dialogVisible = true
    },
    openEditDialog: (index, row) => {
        edit.model.id = row.id
        edit.model.name = row.name
        edit.model.address = row.address
        edit.model.web = row.web
        edit.model.linkman = row.linkman
        edit.model.phone = row.phone
        edit.model.wechatID = row.wechatID
        edit.model.email = row.email
        edit.model.description = row.description
        edit.model.remark = row.remark
        edit.dialogVisible = true
    },
})

const add = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        name: "",
        address: "",
        web: "",
        linkman: "",
        phone: "",
        wechatID: "",
        email: "",
        description: "",
        remark: "",
    },
    submit: () => {
        addForm.value.validate((valid) => {
            if (valid) {
                add.submitDisabled = true
                addSupplier(add.model).then((res) => {
                    if (res.status == 1) {
                        messageForCRUD(add.model.name, "添加成功", "success")
                        base.query()
                    } else {
                        messageForCRUD(add.model.name, "添加失败", "error")
                    }
                    add.dialogVisible = false
                    add.model = {
                        name: "",
                        address: "",
                        web: "",
                        linkman: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                        description: "",
                        remark: "",
                    }
                    add.submitDisabled = false
                })
            } else {
                return false;
            }
        })
    }
})

const view = reactive({
    dialogVisible: false,
    model: {
        name: "",
        address: "",
        web: "",
        linkman: "",
        phone: "",
        wechatID: "",
        email: "",
        description: "",
        remark: "",
    },
})

const edit = reactive({
    dialogVisible: false,
    submitDisabled: false,
    model: {
        id: 0,
        name: "",
        address: "",
        web: "",
        linkman: "",
        phone: "",
        wechatID: "",
        email: "",
        description: "",
        remark: "",
    },
    submit: () => {
        editForm.value.validate((valid) => {
            if (valid) {
                edit.submitDisabled = true
                editSupplier(edit.model).then((res) => {
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
                        address: "",
                        web: "",
                        linkman: "",
                        phone: "",
                        wechatID: "",
                        email: "",
                        description: "",
                        remark: "",
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