<template>
    <el-row>
        <el-col :span="4" :offset="6">
            原位合同总量：{{ base.model.ywTargetLoad }}
        </el-col>
        <el-col :span="4">
            自研合同总量：{{ base.model.zyTargetLoad }}
        </el-col>
        <el-col :span="4">
            渠道合同总量：{{ base.model.qdTargetLoad }}
        </el-col>
    </el-row>
    <el-row style="margin-top: 15px;"></el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :allShow="true" />
</template>

<script setup>
import { reactive, onBeforeMount } from 'vue'
import { queryTopList } from "@/api/my"

import divTable from '@/components/divTable/index.vue'

const base = reactive({
    model: {
        ywTargetLoad: 0,
        zyTargetLoad: 0,
        qdTargetLoad: 0,
    },
    column: {
        headers: [
            {
                prop: "id",
                label: "排名",
            },
            {
                prop: "name",
                label: "名称",
            },
            {
                prop: "finalPercentages",
                label: "完成百分比",
            },
            {
                prop: "targetLoad",
                label: "完成量",
            },
            {
                prop: "taskLoad",
                label: "目标量",
            },
            {
                prop: "notPayment",
                label: "可回款量",
            },
            {
                prop: "ywTargetLoad",
                label: "原位合同量",
            },
            {
                prop: "zyTargetLoad",
                label: "自研合同量",
            },
            {
                prop: "qdTargetLoad",
                label: "渠道合同量",
            },
        ],
        rowStyle: ({ row }) => {
            if (row.finalPercentages > 100) {
                return {
                    //绿色
                    backgroundColor: '#32CD32',
                    color: '#000',
                }
            } else {
                return {
                    //红色
                    backgroundColor: '#FF4500',
                    color: '#000'
                }
            }
            //橙色
            // #FF8C00
        },
    },
    tableData: [],
    query: () => {
        queryTopList().then((res) => {
            if (res.status == 1) {
                base.model = res.data.office
                base.tableData = res.data.offices
            }
        })
    },
})

onBeforeMount(() => {
    base.query()
})
</script>