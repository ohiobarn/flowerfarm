const supressNotices = () => {
	const noticePush = (e) => {
		e.style.display = 'none';
		if (e.classList.contains('hidden')) {
			return;
		}

		let types = [
			'notice-success',
			'notice-error',
			'notice-info',
			'notice-warning',
			'error',
		]

		let type = 'notice-info';

		types.map(eType => type = e.classList.contains(eType) ? eType : type)

		notices.push({
			type: type,
			id: e.getAttribute('id'),
			message: e.innerText,
		});
	}
	let notices = [];
	[...document.querySelectorAll('.notice')].map(e => noticePush(e));
	[...document.querySelectorAll('.error')].map(e => noticePush(e));
	return notices;
}

export default supressNotices;
