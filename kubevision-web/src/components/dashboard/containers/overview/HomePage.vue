<template>
  <v-row>
    <v-col cols="4">
      <card-cluster :cluster="clusterInfo"/>
    </v-col>
  </v-row>
</template>

<script setup>
import { ref } from 'vue';
import API from '@/assets/app/api';

import CardCluster from '@/components/plugins/CardCluster.vue';

var clusterInfo = ref({});

async function refresh() {
  clusterInfo.value = (await API.cluster.get()).cluster
}

refresh();

</script>