<template>
  <div class="app-wrapper">
    <el-container>
      <el-header ref="header" class="app-header">
        <div class="logo">easygen</div>
        <div style="float: left">
          <el-button v-show="currCode != null" class="copy_code" title="复制代码" circle :data-clipboard-text="currCode && currCode.source">
            <el-icon style="vertical-align: middle">
              <Document />
            </el-icon>
          </el-button>
        </div>
        <div style="float: right">
          <el-button style="margin-left: 10px" type="primary" @click="showGenDialog()">生成代码</el-button>
        </div>
      </el-header>
      <el-container>
        <el-aside class="main-left">
          <el-row class="toolbox">
            <el-col :span="14" style="font-size: 14px; font-weight: 500">{{ store.provider.database }}</el-col>
            <el-col :span="10" style="text-align: right; padding-right: 10px">
              <el-button size="small" title="切换数据源" @click="showProviderDialog">
                <el-icon>
                  <Edit />
                </el-icon>
              </el-button>
              <el-button size="small" title="重新加载表" @click="reloadTable">
                <el-icon>
                  <Refresh />
                </el-icon>
              </el-button>
            </el-col>
          </el-row>
          <el-tree class="tree" :data="tables" v-loading="loading" node-key="id" @node-click="selectTable" highlight-current default-expand-all>
          </el-tree>
        </el-aside>
        <el-main class="main-right">
          <el-tabs v-if="codes && codes.length > 0" v-model="codeName" type="border-card" :closable="false" @tab-remove="removeTab">
            <el-tab-pane v-for="(item, index) in codes" :key="index" :label="item.name" :name="item.name">
              <highlightjs :language="item.lang" :autodetect="false" :code="item.code"></highlightjs>
            </el-tab-pane>
          </el-tabs>
        </el-main>
      </el-container>
    </el-container>
    <provider-dialog ref="providerDialogRef"></provider-dialog>
    <gen-dialog ref="genDialogRef" @gen-handle="genHandle"></gen-dialog>
  </div>
</template>

<script setup>
import '@/styles/style.scss'
import { ref, computed, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { Edit, Tools, Refresh, Document, CollectionTag } from '@element-plus/icons-vue'
import { useStore } from '@/store/index.js'
import ui from '@/utils/ui.js'
import Clipboard from 'clipboard'
import ProviderDialog from '@/components/ProviderDialog.vue'
import GenDialog from '@/components/GenDialog.vue'

const providerDialogRef = ref(null)
const genDialogRef = ref(null)
const store = useStore()
const loading = ref(false)

const { tables } = storeToRefs(store)

const table = ref(null)
const codes = ref([])
const codeName = ref(null)

onMounted(async () => {
  await store.init()
  if (!store.isConnected) {
    showProviderDialog()
  }

  initCopy()
})

const initCopy = () => {
  let c = new Clipboard('.copy_code')
  c.on('success', () => {
    ui.success('复制成功')
  })
  c.on('error', () => {
    ui.error('复制失败')
  })
}

const currCode = computed(() => {
  return codes.value.find((e) => e.name == codeName.value)
})

const showProviderDialog = () => {
  providerDialogRef.value.show()
}

const showGenDialog = () => {
  console.log(genDialogRef)
  if (table.value === null) return ui.warning('请选择要生成代码的表')

  genDialogRef.value.show(table.value)
}

const reloadTable = async () => {
  loading.value = true
  await store.getTables()
  table.value = null
  setTimeout(() => {
    loading.value = false
  }, 100)
}

const selectTable = (data) => {
  table.value = data.label
}

const removeTab = (targetName) => {
  let activeName = codeName.value
  if (activeName === targetName) {
    codes.value.forEach((tab, index) => {
      if (tab.name === targetName) {
        const nextTab = codes.value[index + 1] || codes.value[index - 1]
        if (nextTab) {
          activeName = nextTab.name
        }
      }
    })
  }

  codeName.value = activeName
  codes.value = codes.value.filter((tab) => tab.name !== targetName)
}

const genHandle = (data) => {
  codes.value = data.codes
  codeName.value = data.codeName
}
</script>
