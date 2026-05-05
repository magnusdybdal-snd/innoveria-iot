"""
Plots the virtual user ramp profile from Test 3 (concurrent-300.js).
Data source: docs/tests/results/2026-05-02-run-01.md
Output: docs/tests/results/vu_ramp.pdf
"""

import matplotlib.pyplot as plt
import os

# VU ramp — derived from k6 stage config
time_min = [0, 2, 4, 6, 9, 10]
vus      = [0, 300, 600, 1000, 1000, 0]
requirement_vus = 300

fig, ax = plt.subplots(figsize=(7, 4))

ax.fill_between(time_min, vus, alpha=0.15, color='#2E75B6')
ax.plot(time_min, vus, color='#2E75B6', linewidth=2, label='Samtidige brukere (VU)')

ax.axhline(
    y=requirement_vus,
    color='#C00000',
    linestyle='--',
    linewidth=1.5,
    label=f'Krav: {requirement_vus} brukere',
)

# Annotate phases
phases = [
    (1.0,  150,  '0→300'),
    (3.0,  450,  '300→600'),
    (5.0,  800,  '600→1 000'),
    (7.5, 1020,  'Hold'),
    (9.5,  500,  'Ned'),
]
for x, y, label in phases:
    ax.text(x, y, label, ha='center', fontsize=8, color='#404040')

ax.set_xlabel('Tid (minutter)')
ax.set_ylabel('Virtuelle brukere')
ax.set_title('VU-rampe – HTTP belastningstest')
ax.legend(loc='upper left')
ax.grid(axis='y', linestyle=':', alpha=0.4)
ax.set_xlim(0, 10)
ax.set_ylim(0, 1150)
ax.set_xticks(range(0, 11))

plt.tight_layout()

out_dir = os.path.join(os.path.dirname(__file__), '..', 'results')
plt.savefig(os.path.join(out_dir, 'vu_ramp.pdf'), format='pdf', bbox_inches='tight')
plt.savefig(os.path.join(out_dir, 'vu_ramp.png'), dpi=150, bbox_inches='tight')
print("Saved vu_ramp.pdf + .png")
