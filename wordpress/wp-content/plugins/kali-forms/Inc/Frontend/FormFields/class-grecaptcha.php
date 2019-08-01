<?php
namespace KaliForms\Inc\Frontend\FormFields;

if (!defined('ABSPATH')) {
    exit;
}

/**
 * Class Google Recaptcha
 *
 * @package Inc\Frontend\FormFields;
 */
class GRecaptcha extends Form_Field
{
    /**
     * Class constructor
     */
    public function __construct()
    {
        $this->id = 'grecaptcha';
    }

    /**
     * Render function
     *
     * @return void
     */
    public function render($item, $form_info)
    {
        if ($form_info['remove_captcha_for_logged_users'] && is_user_logged_in()) {
            return;
        }

        $div = '<div class="col-12 col-md-' . absint($item['col']) . '">';
        $div .= '<div id="' . $item['id'] . '" data-field-type="grecaptcha" data-sitekey="' . $form_info['google_site_key'] . '"></div>';
        $div .= '</div>';

        return $div;
    }
}
