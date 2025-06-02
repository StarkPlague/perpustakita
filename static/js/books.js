let allBooks = [];

function fetchAndRenderBooks() {
    fetch('/books')
        .then(response => response.json())
        .then(data => {
            allBooks = data;
            renderBookTable(allBooks);
        })
        .catch(err => console.error("Error fetching books:", err));
}

function renderBookTable(data) {
    const tableBody = document.getElementById("book-body");
    tableBody.innerHTML = "";

    data.forEach(book => {
        const row = document.createElement("tr");
        row.innerHTML = `
            <td>${book.title}</td>
            <td>${book.author}</td>
            <td>${book.quantity}</td>
            <td><button onclick="deleteBook(${book.id})">Hapus</button></td>
            <td><button onclick="editBook(${book.id}, '${book.title}', '${book.author}', ${book.quantity})">Edit</button></td>
        `;
        tableBody.appendChild(row);
    });
}

function applySearchandSort() {
    const searchTerm = document.getElementById("searchInput").value.toLowerCase();
    const sortOption = document.getElementById("sortSelect").value;

    let filtered = allBooks.filter(book =>
        book.title.toLowerCase().includes(searchTerm) ||
        book.author.toLowerCase().includes(searchTerm)
    );

    if (sortOption) {
        const [key, order] = sortOption.split("-");
        filtered.sort((a, b) => {
            if (key === "quantity") {
                return order === "asc" ? a.quantity - b.quantity : b.quantity - a.quantity;
            } else {
                return order === "asc"
                    ? a[key].localeCompare(b[key])
                    : b[key].localeCompare(a[key]);
            }
        });
    }

    renderBookTable(filtered);
}

window.onload = function () {
    fetchAndRenderBooks();

    // attach event listener kalau belum ada
    document.getElementById("searchInput").addEventListener("input", applySearchandSort);
    document.getElementById("sortSelect").addEventListener("change", applySearchandSort);
};
