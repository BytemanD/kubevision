<template>
    <v-card color="info" density="compact" append-icon="mdi-kubernetes"
        :title="cluster && cluster.current_context">
        <v-card-subtitle>{{ cluster && cluster.host }} </v-card-subtitle>
        <v-card-text>
            <v-table density="compact">
                <tbody>
                    <tr>
                        <td><strong>git_version</strong></td>
                        <td>{{ cluster.server_version && cluster.server_version.gitVersion }}</td>
                    </tr>
                    <tr>
                        <td><strong>build_date</strong></td>
                        <td>{{ cluster.server_version && cluster.server_version.buildDate }}</td>
                    </tr>
                    <tr>
                        <td><strong>go_version</strong></td>
                        <td>{{ cluster.server_version && cluster.server_version.goVersion }}</td>
                    </tr>
                    <tr>
                        <td><strong>platform</strong></td>
                        <td>{{ cluster.server_version && cluster.server_version.platform }}</td>
                    </tr>
                </tbody>
            </v-table>
        </v-card-text>
    </v-card>
</template>

<script setup>
import SETTINGS from '@/assets/app/settings';
import { useTheme } from 'vuetify'

const progs = defineProps({
    cluster: { type: Object, required: true },
})


const theme = useTheme()

let itemThemeDark = 'themeDark';

function onclickCallback() {
    theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark';
    SETTINGS.ui.setItem(itemThemeDark, theme.global.current.value.dark);
    SETTINGS.ui.save(itemThemeDark);
}

if (SETTINGS.ui.getItem(itemThemeDark).getValue()) {
    theme.global.name.value = 'dark';
} else {
    theme.global.name.value = 'light';
}

</script>