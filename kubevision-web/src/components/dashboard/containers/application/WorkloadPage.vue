<template>
    <tab-windows :tabs="tabs">
      <template v-slot:window-items>
        <v-window-item><deployment-page /></v-window-item>
        <v-window-item><daemonset-page /></v-window-item>
        <v-window-item><statefulset-page /></v-window-item>
        <v-window-item><job-page /></v-window-item>
        <v-window-item><cronjob-page /></v-window-item>
      </template>
    </tab-windows>
</template>

<script setup>
import { ref } from 'vue';

import TabWindows from '@/components/plugins/TabWindows.vue';
import DeploymentPage from '@/components/dashboard/containers/application/DeploymentPage.vue';
import DaemonsetPage from '@/components/dashboard/containers/application/DaemonsetPage.vue';
import StatefulsetPage from '@/components/dashboard/containers/application/StatefulsetPage.vue';
import JobPage from '@/components/dashboard/containers/application/JobPage.vue';
import CronjobPage from '@/components/dashboard/containers/application/CronjobPage.vue';

var tabs = ref([
  {text: 'deployment'},
  {text: 'daemonset'},
  {text: 'statefulset'},
  {text: 'job'},
  {text: 'cronjob'},
])

</script>