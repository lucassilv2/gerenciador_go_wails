<script>
  import { ListProcesses, KillProcess, PauseProcess, ResumeProcess } from '../wailsjs/go/main/App.js';
  import { onMount, onDestroy } from 'svelte';

  let processos = [];
  let interval = null;

  function stateTranslate(state) {
    const base = state.charAt(0);
    let stat = '';

    switch (base) {
      case 'S': stat = 'Sleeping'; break;
      case 'R': stat = 'Running'; break;
      case 'D': stat = 'Uninterruptible Sleep'; break;
      case 'T': stat = 'Stopped'; break;
      case 'Z': stat = 'Zombie'; break;
      case 'I': stat = 'Idle'; break;
      default: return 'Unknown';
    }

    if (state.includes('s')) stat += ' (Session Leader)';
    if (state.includes('+')) stat += ' (Foreground)';
    if (state.includes('N')) stat += ' (Low Priority)';
    if (state.includes('<')) stat += ' (High Priority)';
    if (state.includes('l')) stat += ' (Multithreaded)';

    return stat;
  }

  function parseProcessList(lines) {
    const [_, ...data] = lines.filter(line => line.trim() !== '');

    return data.map(line => {
      const pid = parseInt(line.slice(0, 6).trim());
      const user = line.slice(6, 22).trim();
      const cpu = parseFloat(line.slice(22, 30).trim().replace(',', '.'));
      const time = line.slice(30, 42).trim();
      const command = line.slice(42, line.length - 5).trim();
      const stat = line.slice(-5).trim();

      if (isNaN(pid) || isNaN(cpu)) return null;

      return { pid, user, cpu, time, command, stat };
    }).filter(Boolean); // remove nulls
  }

  async function loadProcesses() {
    try {
      const result = await ListProcesses(); // já é um array
      processos = parseProcessList(result);
    } catch (err) {
      console.error("Erro ao carregar processos:", err);
    }
  }

  function kill(pid) {
    KillProcess(pid.toString()).then(loadProcesses);
  }

  function pause(pid) {
    PauseProcess(pid.toString()).then(loadProcesses);
  }

  function resume(pid) {
    ResumeProcess(pid.toString()).then(loadProcesses);
  }

  onMount(() => {
    loadProcesses();
    interval = setInterval(loadProcesses, 1000);
    
  });

  onDestroy(() => {
    clearInterval(interval);
  });
</script>

<main>
  <div style="flex: 1; overflow: auto;">
    <table>
      <thead>
        <tr>
          <th>PID</th>
          <th>Usuário</th>
          <th>%CPU</th>
          <th>Tempo</th>
          <th>Comando</th>
          <th>Estado</th>
          <th>Ações</th>
        </tr>
      </thead>
      <tbody>
        {#each processos as proc}
          <tr>
            <td>{proc.pid}</td>
            <td>{proc.user}</td>
            <td><span style="color: {proc.cpu > 10 ? '#ff5555' : 'inherit'}">{proc.cpu.toFixed(1)}</span></td>
            <td>{proc.time}</td>
            <td>{proc.command}</td>
            <td title={proc.stat}>{stateTranslate(proc.stat)}</td>
            <td>
              {#if proc.stat.includes('T')}
                <button class="resume" on:click={() => resume(proc.pid)}>▶️</button>
              {:else}
                <button class="pause" on:click={() => pause(proc.pid)}>⏸️</button>
              {/if}
              <button class="kill" on:click={() => kill(proc.pid)}>🗑️</button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</main>
