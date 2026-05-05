"""
Plots HTTP response time percentiles from Test 3 (concurrent-300.js).
Data source: docs/tests/results/2026-05-02-run-01.md
Output: docs/tests/results/response_time_percentiles.pdf
"""

import matplotlib.pyplot as plt
import matplotlib.ticker as ticker
import os

# Data from Test 3 — peak 1 000 VUs
labels = ['p50', 'p90', 'p95', 'Maks']
values = [0.96, 7.20, 8.68, 68.14]
requirement_ms = 1000

fig, ax = plt.subplots(figsize=(7, 4.5))

colors = ['#5B9BD5', '#5B9BD5', '#2E75B6', '#1F4E79']
bars = ax.bar(labels, values, color=colors, width=0.5, zorder=3)

ax.set_ylim(0, 90)
ax.set_ylabel('Svartid (ms)')
ax.set_xlabel('Persentil')
ax.set_title('HTTP svartid ved topp 1 000 samtidige brukere')
ax.grid(axis='y', linestyle=':', alpha=0.4, zorder=0)

# Text box referencing requirement without a line symbol
ax.text(
    0.02, 0.97,
    f'Krav: {requirement_ms} ms',
    transform=ax.transAxes,
    fontsize=9,
    verticalalignment='top',
    bbox=dict(boxstyle='round,pad=0.4', facecolor='white', edgecolor='#AAAAAA', linewidth=0.8),
)

for bar, val in zip(bars, values):
    ax.text(
        bar.get_x() + bar.get_width() / 2,
        val + 1.2,
        f'{val} ms',
        ha='center',
        va='bottom',
        fontsize=9,
    )

plt.tight_layout()

out_dir = os.path.join(os.path.dirname(__file__), '..', 'results')
plt.savefig(os.path.join(out_dir, 'response_time_percentiles.pdf'), format='pdf', bbox_inches='tight')
plt.savefig(os.path.join(out_dir, 'response_time_percentiles.png'), dpi=150, bbox_inches='tight')
print("Saved response_time_percentiles.pdf + .png")
