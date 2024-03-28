<template>
    <div>
        <v-chip :label="label" :small="small" :x-small="xSmall" :large="large" :x-range="xLarge" class="mr-0 no-right-radius"
            :color="nameColor" >{{ name }}</v-chip>
        <v-chip :small="small" :x-small="xSmall" :large="large" :x-range="xLarge" class="ml-0 no-left-radius"
            :color="valueColor" >{{ value }}</v-chip>
    </div>
</template>

<style>
.no-left-radius {
    border-top-left-radius: 0 !important;
    border-bottom-left-radius: 0 !important;
}
.no-right-radius {
    border-top-right-radius: 0 !important;
    border-bottom-right-radius: 0 !important;
}
</style>

<script>

export default {
    props: {
        name: String,
        value: [String, Number],
        nameColor: {
            type: String, default: 'primary'
        },
        label: { type: Boolean, default: false },
        valueColor: { type: String, default: 'info' },
        small: { type: Boolean, default: false },
        xSmall: { type: Boolean, default: false },
        large: { type: Boolean, default: false },
        xLarge: { type: Boolean, default: false },
    },

};


</script>