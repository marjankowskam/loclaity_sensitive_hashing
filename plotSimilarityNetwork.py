import pandas as pd
import networkx as nx
import matplotlib.pyplot as plt
import plotly.graph_objects as go

# Load the weighted adjacency matrix from a file
filename = 'data/similarity_array_network.csv'
adj_matrix = pd.read_csv(filename, header=None).values

with open('data/similarity_array_network_labels.csv', 'r') as file:
    node_labels = file.readline().strip().split(',')

# Create a graph from the adjacency matrix
G = nx.from_numpy_array(adj_matrix)

# Remove self-loops and edges with weights less than cutoff
cutoff = 0.3
G.remove_edges_from([(u, v) for u, v, w in G.edges(data=True) if u == v or w['weight'] < cutoff])

# Extract edge weights
weights = nx.get_edge_attributes(G, 'weight')

# Plot the graph
pos = nx.spring_layout(G)  # positions for all nodes


# Normalize edge weights for color mapping
min_weight = min(weights.values())
max_weight = max(weights.values())

# Function to map weights to grayscale colors
def weight_to_color(weight, min_weight, max_weight):
    norm_weight = (weight - min_weight) / (max_weight - min_weight)
    gray_value = int(255 * (1-norm_weight))
    return f'rgba({gray_value}, {gray_value}, {gray_value}, 0.8)'

# # Draw nodes and edges
# nx.draw_networkx_nodes(G, pos, node_size=300, alpha=0.5)
# nx.draw_networkx_edges(G, pos, width=1.0, alpha=0.5)

# # Draw node and edge labels
#nx.draw_networkx_edge_labels(G, pos, edge_labels=weights, font_size=10)
#nx.draw_networkx_labels(G, pos, font_size=5, font_family='sans-serif')


# Create edge traces
edge_trace = []
eweigths_trace = []; xtext=[]; ytext=[]; hover_texts = []
for edge in G.edges(data=True):
    x0, y0 = pos[edge[0]]
    x1, y1 = pos[edge[1]]
    weight = edge[2]['weight']
    color = weight_to_color(weight, min_weight, max_weight)
    edge_trace.append(go.Scatter(
        x=[x0, x1, None],
        y=[y0, y1, None],
        line=dict(width=2, color=color),
        mode='lines',
        hoverinfo='none'
    ))
    xtext.append((x0+x1)/2)
    ytext.append((y0+y1)/2)
    hover_texts.append(f'Weight: {weight:.2f}')
    
eweights_trace = go.Scatter(
    x=xtext,
    y= ytext, 
    mode='markers',
    text=hover_texts,
    textposition='top center',
    hoverinfo='text',    
    marker=dict(size=1, color='rgba(0,0,0,0)')
)
# Create node trace
node_trace = go.Scatter(
    x=[],
    y=[],
    text=[],
    mode='markers',
    textposition='top center',
    hoverinfo='text',
    marker=dict(
        showscale=False,
        colorscale='YlGnBu',
        size=10,
        line_width=2))

for node in G.nodes():
    x, y = pos[node]
    node_trace['x'] += (x,)
    node_trace['y'] += (y,)
    node_trace['text'] += (node_labels[node],)

# Create the figure
fig = go.Figure(data=edge_trace + [eweights_trace, node_trace],
                layout=go.Layout(
                    title='Clustering of Hostnames based on LSH (edge cutoff 0.3)',
                    titlefont_size=16,
                    showlegend=False,
                    hovermode='closest',
                    margin=dict(b=20, l=5, r=5, t=40),
                    annotations=[dict(
                        text="Node labels appear on hover",
                        showarrow=False,
                        xref="paper", yref="paper",
                        x=0.005, y=-0.002)],
                    xaxis=dict(showgrid=False, zeroline=False),
                    yaxis=dict(showgrid=False, zeroline=False)))

# Display the plot
fig.show()





