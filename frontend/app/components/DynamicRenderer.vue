<script setup>
defineProps({
  data: {
    type: [Object, Array, String, Number, Boolean],
    required: true
  }
})
</script>

<template>
  <div v-for="(value, key) in data">
    <div v-if="typeof value !== 'object' && !Array.isArray(value)"><strong style="color:yellow">{{ key }}</strong>: {{
      value }}</div>
    <div v-else> <strong style="color:brown">{{ key }}</strong></div>
    <DynamicRenderer
      v-if="typeof value === 'object' && !Array.isArray(value)"
      :data="value"
    />
    <div v-else-if="Array.isArray(value)">
      <UTable :data="value" />
    </div>
  </div>
</template>
