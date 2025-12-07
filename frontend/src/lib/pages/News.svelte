<script lang="ts">
  import CardFeed from '../components/CardFeed.svelte';

  const newsItems = [
    {
      title: 'Winter Event Coming Soon',
      date: 'Dec 5, 2025',
      featured: true,
      content: 'Prepare for snowfall, exclusive rewards, and limited-time challenges.'
    },
    {
      title: 'Server Maintenance Complete',
      date: 'Dec 3, 2025',
      content: 'All servers are back online. Thank you for your patience.'
    }
  ];
</script>

<CardFeed title="News & Events" items={newsItems} />