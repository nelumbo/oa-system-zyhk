<template>
    <!-- <el-row>
        <el-col :span="4" :offset="6">
            原位合同总量：{{ base.model.ywTargetLoad }}
        </el-col>
        <el-col :span="4">
            自研合同总量：{{ base.model.zyTargetLoad }}
        </el-col>
        <el-col :span="4">
            渠道合同总量：{{ base.model.qdTargetLoad }}
        </el-col>
    </el-row> -->
    <el-row style="margin-top: 15px;"></el-row>
    <divTable :columnObj="base.column" :tableData="base.tableData" :allShow="true" />
    <div style="margin-top: 30px;">
        <el-divider content-position="left">
            <h2>技术排行榜</h2>
        </el-divider>
        <divTable :columnObj="base.column2" :tableData="base.tableData2" :allShow="true" />
    </div>
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
        showIndex: true,
        rowStyle: ({ row }) => {
            if (row.finalPercentages >= 100) {
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
    column2: {
        headers: [
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
            }
        ],
        showIndex: true,
        rowStyle: ({ row }) => {
            if (row.finalPercentages >= 100) {
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
    tableData2: [],
    query: () => {
        queryTopList().then((res) => {
            if (res.status == 1) {
                base.model = res.data.office
                base.tableData = res.data.offices
                base.tableData2 = res.data.offices_js
            }
        })
    },
})

onBeforeMount(() => {
    base.query()
})
</script>